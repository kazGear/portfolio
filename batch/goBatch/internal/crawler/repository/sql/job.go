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
             ;
    `
}

func InsertJob() string {
	return `
        INSERT INTO t_jobs
        (
            title,
            url,
            description,
            created_at
        )
            VALUES
        (
            :title,
            :url,
            :description,
            NOW()
        );
    `
}