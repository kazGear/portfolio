package sql

func UpdateJob() string {
	return `
        UPDATE
               t_jobs
           SET
               is_active = :is_active,
               updated_at = NOW()
         WHERE
               url = :url
           AND (is_active = TRUE OR is_active IS NULL)
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
            min_salary_at_hour,
            max_salary_at_hour,
            min_salary_at_month,
            max_salary_at_month,
            skills_text,
            required_skills_text,
            preferred_skills_text,
            description,
            employment_type,
            remote_type,
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
            :min_salary_at_hour,
            :max_salary_at_hour,
            :min_salary_at_month,
            :max_salary_at_month,
            :skills_text,
            :required_skills_text,
            :preferred_skills_text,
            :description,
            :employment_type,
            :remote_type,
            :is_active,
            :similarity_score,
            :source_site,
            NOW()
        );
    `
}