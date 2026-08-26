package sql

func UpdateJob() string {
	return `
        UPDATE
               t_jobs
           SET
               title               = :title,
               location            = :location,
               min_salary_at_month = :min_salary_at_month,
               max_salary_at_month = :max_salary_at_month,
               employment_type     = :employment_type,
               work_place          = :work_place,
               updated_at          = :updated_at
         WHERE
               url = :url
             ;
    `
}

func InsertJob() string {
	return `
        INSERT INTO t_jobs
        (
            title,
            url,
            location,
            min_salary_at_month,
            max_salary_at_month,
            description,
            employment_type,
            work_place,
            source_site,
            created_at,
            updated_at
        )
        VALUES
        (
            :title,
            :url,
            :location,
            :min_salary_at_month,
            :max_salary_at_month,
            :description,
            :employment_type,
            :work_place,
            :source_site,
            NOW(),
            :updated_at
        );
    `
}

func SelectCreatedFeatures() string {
	return `
        SELECT
               job_id
          FROM
               t_job_features_created
             ;
    `
}

func SelectCurrentJobId() string {
	return `
        SELECT
               id
          FROM
               t_jobs
         WHERE
               url = :url
             ;
    `
}

func InsertJobId() string {
	return `
        INSERT INTO t_job_features_created (job_id) VALUES ($1);
    `
}

func SelectSavedPageIds() string {
	return `
        SELECT
              (regexp_match(url, '\d{1,7}'))[1]::INT
          FROM
               t_jobs
         WHERE
               source_site = $1
             ;
    `
}