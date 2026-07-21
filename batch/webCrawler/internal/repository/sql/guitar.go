package sql

func UpdateGuitar() string {
	return `
        UPDATE t_guitars
           SET body_finish         = :body_finish,
               body_material       = :body_material,
               body_material_top   = :body_material_top,
               body_material_back  = :body_material_back,
               bridge              = :bridge,
               color               = :color,
               color_cd            = :color_cd,
               controls            = :controls,
               comment             = :comment,
               fingerboard         = :fingerboard,
               fret_count          = :fret_count,
               inlays              = :inlays,
               joint               = :joint,
               neck_material       = :neck_material,
               pickups             = :pickups,
               price               = :price,
               scale_length_mm     = :scale_length_mm,
               series              = :series,
               src                 = :src,
               weight              = :weight,
               updated             = NOW()
         WHERE maker = :maker
           AND name  = :name
           AND color = :color
    `
}

func InsertGuitar() string {
	return `
		INSERT INTO t_guitars
		(
			maker,
			name,
			body_finish,
			body_material,
			body_material_top,
			body_material_back,
			bridge,
			color,
			color_cd,
			controls,
			comment,
			fingerboard,
			fret_count,
			inlays,
			joint,
			neck_material,
			pickups,
			price,
			scale_length_mm,
			series,
			src,
			weight,
			updated
		)
			VALUES
		(
			:maker,
			:name,
			:body_finish,
			:body_material,
			:body_material_top,
			:body_material_back,
			:bridge,
			:color,
			:color_cd,
			:controls,
			:comment,
			:fingerboard,
			:fret_count,
			:inlays,
			:joint,
			:neck_material,
			:pickups,
			:price,
			:scale_length_mm,
			:series,
			:src,
			:weight,
			NOW()
		)
	`
}