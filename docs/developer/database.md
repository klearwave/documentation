# Database Standards

This page contains the agreed upon developer standards as it pertains to
databases.


## Design Standards


### Migrations: Goose

[Goose](https://github.com/pressly/goose) should be used in conjunction 
with SQL files in order to achieve database migrations.  Migrations should exist at the top level
of a repository in a `migrations/` folder.

> Example `migrations/` directory:

```bash
tree service-info/migrations/
service-info/migrations/
├── 000001_create_versions_table.sql
└── 000002_create_container_images_table.sql
```

> **Example `migration.sql` file**:

```sql
-- +goose Up
CREATE TABLE versions (
    id SERIAL PRIMARY KEY,
    version_id VARCHAR(32) NOT NULL UNIQUE,
    stable BOOLEAN DEFAULT FALSE,
    x_version INT,
    y_version INT,
    z_version INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
);

-- +goose Down
DROP TABLE versions;
```


### Code: GORM

[GORM](https://gorm.io/docs/) should be used in conjunction with Go structs
in order to tie database models to their underlying records:

> **Example struct with GORM**:

```go
// ContainerImageBase is the base set of fields for all ContainerImage objects.
type ContainerImageBase struct {
	Image      string `json:"image,omitempty" example:"ghcr.io/klearwave/service-info:latest" doc:"Full container image including the registry, repository and tag."`
	SHA256Sum  string `json:"sha256sum,omitempty" example:"2d4b92db6941294f731cfe7aeca336eb8dba279171c0e6ceda32b9f018f8429d" doc:"SHA256 sum of the container image."`
	CommitHash string `json:"commit_hash,omitempty" example:"631af50a8bbc4b5e69dab77d51a3a1733550fe8d" doc:"Commit hash related to the image."`
}
```


### Code: Validation and Mutation

Use Golang as opposed to writing complex functions 
in SQL, we should use proper code to validate inputs before inserting into the 
database.  Additionally, we should also use proper code to mutate any incoming 
database records.

> **Example (Do not do this)**:

Here is a sample SQL file with complex trigger functions (logic in SQL):

```sql
-- Function to validate semantic versioning format
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION trigger_validate_version_id()
RETURNS TRIGGER AS $$
BEGIN
    IF NOT NEW.version_id ~ '^v\d+\.\d+\.\d+$' THEN
        RAISE EXCEPTION 'Invalid version_id format. Must be in the form vX.Y.Z';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
```

> **Example (Do this instead)**:

Here is the same logic, written in code:

```go
// ValidateVersionID checks if the given version ID follows the pattern vX.Y.Z
func ValidateVersionID(versionID string) error {
	validVersionPattern := `^v\d+\.\d+\.\d+$`
	re := regexp.MustCompile(validVersionPattern)
	if !re.MatchString(versionID) {
		return errors.New("Invalid version_id format. Must be in the form vX.Y.Z")
	}
	return nil
}
```

The above shows an example of a trigger function in SQL for validation.  While this does 
enforce the data inserted closest to where it lives, it creates unneeded complexity and 
additional technical debt (need to understand writing of SQL functions and the codebase 
that calls it) that we want to avoid.  Additionally, services that run proper code are
much easier to scale horizontally, so it is better to use resources outside of the database 
in proper code than to stuff it all in the database.

However, there is are exceptions to the rule...


#### Code: Validation Exceptions

1. **Length Validation**: In the instance of length validation, it is better to enforce length validation at the 
database level to avoid the need to validate in code and to enforce the data integrity 
of what is stored in the database.  This is specific to string data which can be of 
variable length.  We should look to implement string length validation in one of two manners:

* **Variable Length**: for variable length strings (strings which do not have a predetermined length), we 
should use the `VARCHAR(n)` column type:

```sql
    version_id VARCHAR(32) NOT NULL UNIQUE,
```

In the above example, a `version_id` can be different lengths (e.g. `v0.1.0` versue `v0.1.10`), so we 
should use the `VARCHAR(n)` type.  This allows the database to not preallocate space and validates 
that the data only allows a specific length (32).

* **Non-Variable Length**: for string data that has a pre-determined length, we should use the 
`CHAR(n)` column type:

```sql
    sha256_sum CHAR(64) NOT NULL,
    commit_hash CHAR(40) NOT NULL,
```

In the above example, a `sha256_sum` and `commit_hash` have pre-determined lengths and will always 
only be those lengths, so we use the `CHAR(n)` to preallocate the space.
