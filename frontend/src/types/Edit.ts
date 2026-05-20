export interface EditMonsterDTO {
    MonsterId: string;
    MonsterName: string;
    MonsterType: number;
    MonsterTypeName: string;
    Hp: number;
    Attack: number;
    Speed: number;
    Week: number;
    WeekName: string;
    // 変更後パラメータ
    AfterName: string | undefined;
    AfterHp: number | undefined;
    AfterAttack: number | undefined;
    AfterSpeed: number | undefined;
    AfterWeek: number |  undefined;
    IsChanged: boolean;
}

export interface EditSkillsDTO {
    // モンスターステータス
    ItemId: string;
    MonsterId: string;
    MonsterName: string;
    Hp: number;
    MonsterAttack: number;
    Speed: number;
    WeekName: string;
    MySkillId: string;
    SkillId: string;
    SkillName: string;
    SkillAttack: number;
    SkillElementName: string;
    IsChanged: boolean;

    // 各スキル
    MySkillIds: string[];
    SkillIds: string[];
    SkillNames: string[];
    SkillAttacks: number[];
    SkillElementNames: string[];
}

export interface AllSkillDTO {
    SkillId: string;
    SkillName: string;
    SkillTypeName: string;
    ElementName: string;
    Attack: number;
}