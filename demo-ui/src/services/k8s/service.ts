import { request } from "@umijs/max";

export async function listService(
  params: {
      // query
      /** 当前的页码 */
      current?: number;
      /** 页面的容量 */
      pageSize?: number;
      name?: string;
      namespace?: string;
  },
  options?: { [key: string]: any },
  ) {
  return request<API.ServiceList>(`${API_BASE_URL}/service`, {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
    })
  }

export async function addService(
  body: API.Service,
  options?: { [key: string]: any },
  ) {
  return request(`${API_BASE_URL}/service`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
   ...(options || {}),
    })
}

export async function removeService(
  name: string,
  options?: { [key: string]: any },
  ) {
  return request(`${API_BASE_URL}/service/${name}`, {
    method: 'DELETE',
  ...(options || {}),
    })
}

export async function updateService(
  name: string,
  body: API.Service,
  options?: { [key: string]: any },
  ) {
  return request(`${API_BASE_URL}/service/${name}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
  ...(options || {}),
  })
}