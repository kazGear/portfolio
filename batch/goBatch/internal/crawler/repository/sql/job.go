package sql

func UpdateJob() string {
	return `
        UPDATE
               t_jobs
           SET
               is_active  = :is_active,
               updated_at = NOW()
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
            company_name,
            location,
            min_salary_at_month,
            max_salary_at_month,
            description,
            employment_type,
            work_place,
            is_active,
            similarity_score,
            source_site,
            created_at
        )
        VALUES
        (
            :title,
            :url,
            :company_name,
            :location,
            :min_salary_at_month,
            :max_salary_at_month,
            :description,
            :employment_type,
            :work_place,
            :is_active,
            :similarity_score,
            :source_site,
            NOW()
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