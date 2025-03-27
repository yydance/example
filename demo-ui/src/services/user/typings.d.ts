declare namespace API {
    type UserInfo = {
        username: string;
        email?: string;
        password?: string;
    };
    type UserOutput = {
        id?: number;
        name?: string;
        email?: string;
    };
    type UserList = {
    data?: UserOutput[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
    };
}