import { request } from "@umijs/max";

export async function listUser(
    params: {
        // query
        /** 当前的页码 */
        current?: number;
        /** 页面的容量 */
        pageSize?: number;
      },
    options?: { [key: string]: any },
) {
    return request<API.UserList>(`${API_BASE_URL}/user/list`, {
        method: "GET",
        params: {
            ...params,
        },
       ...(options || {}),
    })
}

export async function addUser(
    body: API.UserInfo,
    options?: { [key: string]: any }) {
    return request<Record<string, any>>(`${API_BASE_URL}/user`, {
        method: "POST",
        data: body,
     ...(options || {}),
    })
}

export async function deleteUser(id: number,
    options?: { [key: string]: any }) {
    return request<Record<string, any>>(`${API_BASE_URL}/user/${id}`, {
        method: "DELETE",
      ...(options || {}),
    })
}

export async function updateUser(id: number,
    options?: { [key: string]: any }) {
    return request<Record<string, any>>(`${API_BASE_URL}/user/${id}`, {
        method: "PUT",
     ...(options || {}),
    })
}