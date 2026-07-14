import { PRIVATE_API_BASE_URL } from "../config/env"

const ACCENT_COLOR: string = "#0aff84";
const ACCENT_COLOR2: string = "#F05F8D";
const ALERT_COLOR: string = "red";
const SHADOW_COLOR: string = "#15400e";

export const COLORS = {
    MAIN_FONT_COLOR: "gray",
    ACCENT_FONT_GREEN: `${ACCENT_COLOR}`,
    ACCENT_FONT_PINK: `${ACCENT_COLOR2}`,
    LOSER_FONT_COLOR: "#EC008C",
    CAPTION_FONT_COLOR: "#33cc99",
    BUTTON_FONT_COLOR: "gray",

    ALERT_MESSAGE_COLOR: `${ALERT_COLOR}`,

    BORDER_COLOR: "darkgray",
    LINE_COLOR: "blue",

    SHADOW: `${SHADOW_COLOR}`,
    DIALOG_SHADOW: `${SHADOW_COLOR}`,
    MENU_SHADOW: `${SHADOW_COLOR}`,

    BUTTON_COLOR: `${ACCENT_COLOR}`,
    BUTTON_DISABLED: 0.3,

    THEME_COLOR: "#7fffd4",
    BASE_BACKGROUND: "white",
    DARK_BACKGROUND: "#454545",
    MODAL_BACKGROUND: "black",

    MENU_DISABLED: "darkgray",

    DIALOG_FRAME: `${ACCENT_COLOR}`,

    LOGINED_COLOR: `${ACCENT_COLOR}`,

    URL: "skyblue"
} as const;

export const GUITAR_COLORS = {
    1: "Red",
    2: "Pink",
    3: "Orange",
    4: "Yellow",
    5: "Green",
    6: "SkyBlue",
    7: "Blue",
    8: "Purple",
    9: "Gray",
    10: "Black",
    11: "White",
    12: "#ceb46c", // natural
    13: "Brown",
    14: "Gold",
    15: "Silver",
    99: "Black", // 例外は黒で表示する
} as const;


export const SIZE = {
    INPUT_WIDTH: "150px",
    INPUT_HEIGHT: "22px",
    HEADER_HEIGHT: "55px",
} as const;

export const STATE_TYPE = {
    NORMAL: 0,
    POISON: 1,
    SLEEP: 2,
    CHARM: 3,
    SLOW: 4,
    POWER_UP: 5,
    DODGE_UP: 6,
    CRITICAL_UP: 7,
    AUTO_HEAL: 8
} as const;

export const STATE_NAME = {
    NORMAL: "正常",
    LOSER: "戦闘不能",
    POISON: "毒",
    SLEEP: "睡眠",
    CHARM: "魅了",
    SLOW: "スロー",
    POWER_UP: "攻撃力UP",
    DODGE_UP: "回避力UP",
    CRITICAL_UP: "クリティカルUP",
    AUTO_HEAL: "自動回復"
} as const;

export const DAMAGE_VIEW = {
    DODGE_START: 0,
    DODGE_END: 1500,
    DAMAGE_END: 2500,
} as const;

export const URLS = {
    // 基本情報取得
    USER_INFO: `${PRIVATE_API_BASE_URL}/api/user/userInfo`,
    MONSTERS_INFO: `${PRIVATE_API_BASE_URL}/api/battle/monstersInfo`,
    ITEM_INFO: `${PRIVATE_API_BASE_URL}/api/shop/itemInfo`,
    FETCH_ELEMENT_CODE: `${PRIVATE_API_BASE_URL}/api/common/FetchElementCode`,
    // ユーザ関係
    REGIST_USER_INIT: `${PRIVATE_API_BASE_URL}/api/user/init`,
    SELECT_LOGIN_USER: `${PRIVATE_API_BASE_URL}/api/user/loginUser`,
    RECORD_USER_RESULT: `${PRIVATE_API_BASE_URL}/api/user/recordUserResults`,
    REGIST_USER: `${PRIVATE_API_BASE_URL}/api/user/userRegist`,
    RESTART_AS_PLAYER: `${PRIVATE_API_BASE_URL}/api/user/restartAsPlayer`,
    GET_MONSTER_COUNT: `${PRIVATE_API_BASE_URL}/api/user/getMonsterCount`,
    // バトル、モンスター関係
    INIT_MONSTERS: `${PRIVATE_API_BASE_URL}/api/battle/init`,
    BET_RATE: `${PRIVATE_API_BASE_URL}/api/battle/betRate`,
    BATTLE_NEXT_TURN: `${PRIVATE_API_BASE_URL}/api/battle/nextTurn`,
    RECORD_BATTLE_RESULT: `${PRIVATE_API_BASE_URL}/api/battle/recordResults`,
    // レポート関係
    INIT_BATTLE_REPORT: `${PRIVATE_API_BASE_URL}/api/battleReport/init`,
    MONSTER_REPORTS: `${PRIVATE_API_BASE_URL}/api/battleReport/monsterReport`,
    BATTLE_REPORTS: `${PRIVATE_API_BASE_URL}/api/battleReport/battleReport`,
    // 認証系
    LOGIN_USER: `${PRIVATE_API_BASE_URL}/api/auth/login`,
    CHECK_LOGIN_TOKEN: `${PRIVATE_API_BASE_URL}/api/auth/checkToken`,
    // ショップ系
    SHOP_INIT: `${PRIVATE_API_BASE_URL}/api/shop/init`,
    SELECT_SHOP_ITEMS: `${PRIVATE_API_BASE_URL}/api/shop/items`,
    PURCHASE_ITEM: `${PRIVATE_API_BASE_URL}/api/shop/purchase`,
    // 設定系
    EDIT_INIT: `${PRIVATE_API_BASE_URL}/api/edit/init`,
    FETCH_EDIT_MONSTERS: `${PRIVATE_API_BASE_URL}/api/edit/fetchMonsters`,
    UPDATE_MONSTER_STATUS: `${PRIVATE_API_BASE_URL}/api/edit/updateMonsterStatus`,
    INIT_ALL_MONSTERS_STATUS: `${PRIVATE_API_BASE_URL}/api/edit/initAllMonsterStatus`,
    INIT_ALL_MONSTERS_SKILLS: `${PRIVATE_API_BASE_URL}/api/edit/initAllMonsterSkills`,
    FETCH_EDIT_SKILLS: `${PRIVATE_API_BASE_URL}/api/edit/FecthEditSkills`,
    FETCH_ALL_SKILLS: `${PRIVATE_API_BASE_URL}/api/edit/fecthAllSkills`,
    CHANGE_MONSTER_SKILLS: `${PRIVATE_API_BASE_URL}/api/edit/UpdateMonsterSkills`,
    // その他
    UPLOAD_IMAGE: `${PRIVATE_API_BASE_URL}/api/common/imgUpload`
} as const;

export const KEYS = {
    TOKEN: "token",
    USER_ID: "userId",
    USER_ROLE: "userRole",
    ORDER_BY_ASC: "ASC",
    ORDER_BY_DESC: "DESC"
} as const;

export const PREFIX = {
    BASE64: "data:image/jpeg;base64,"
} as const;

export const DECO = {
    BLOCK_LINE: "■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋",
    BLOCK_LINE_R: "＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■＋■"
} as const;

export const EDIT_TYPE = {
    MONSTER_STATUS: 1,
    MONSTER_SKILL: 2,
    USE_MONSTER: 3
} as const;

export const USER_ROLE = {
    NORMAL: 1,
    EXCELLENT: 2,
    DUCK: 3,
    BE_CAREFUL: 4,
    BLACK_LIST: 5,
    ADMIN: 90,
    SUPER_ADMIN: 91
} as const;

export const GUITAR = {
    OPEN_PRICE: -1,
    UNDEFINED_PRICE: -2,
    PARSE_ERROR_PRICE: -3,

    INVALID_NUMBER: -1,
    UNkNOWN: 99,
} as const;