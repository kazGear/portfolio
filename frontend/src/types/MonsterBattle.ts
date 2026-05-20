export interface MetaDataDTO {
    TargetMonsterId: string;
    BeforeHp: number;
    ImpactPoint: number;
    StateName: string;
    EnableState: boolean;
    DisableState: boolean;
    SkillId: string;
    EffectTime: number;
    Message: string;
    IsStop: boolean;
    IsDodge: boolean;
    AllLoser: boolean;
    ExistWinner: boolean;
    WinnerMonsterId: string;
    WinnerMonsterName: string;
}

export interface MonsterDTO {
    MonsterId: string;
    MonsterName: string;
    MonsterType: number;
    Hp: number;
    MaxHp: number;
    Attack: number;
    DefaultAttack: number;
    Speed: number;
    DefaultSpeed: number;
    Dodge: number;
    DefaultDodge: number;
    Week: number;
    Skills: SkillDTO[];
    Status: StateDTO[];
    BetScore: number;
    BetRate: number;
}

export interface BattleResults {
    Monsters: MonsterDTO[];
    BattleLog: MetaDataDTO[];
}

export interface SkillDTO {
    SkillId: string;
    SkillName: string;
    SkillType: number;
    ElementType: number;
    StateType: number;
    TargetType: number;
    Attack: number;
    Weight: number;
    DefaultCritical: bigint;
    Critical: bigint;
    EffectTime: number;
    Remarks?: string;
}

export interface StateDTO {
    CodeId: string;
    StateType: number;
    Name: string;
    ShortName: string;
    CancelRate: number;
    Activate: boolean;
}