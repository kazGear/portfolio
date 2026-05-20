export interface MonsterTypeDTO {
    MonsterTypeId: number;
    MonsterTypeName: string;
}

export interface MonsterReportDTO {
    MonsterId: string;
    MonsterName: string;
    BattleCount: number;
    Wins: number;
    WinRate: string;
}

export interface BattleReportDTO {
    BattleId: number;
    BattleEndDate: Date;
    BattleEndDateStr: string;
    BattleEndTime: string;
    BattleEndTimeStr: string;
    Serial: number;
    MonsterId: string;
    MonsterName: string;
    IsWin: boolean;
}