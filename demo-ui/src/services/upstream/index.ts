import { request } from "@umijs/max";

export async function listUpstream(
  params: {
      // query
      /** 当前的页码 */
      current?: number;
      /** 页面的容量 */
      pageSize?: number;
      name?: string;
      type?: string;
  },
  options?: { [key: string]: any },
  ) {
  return request<API.UpstreamList>(`${API_BASE_URL}/upstream`, {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
    })
}

export async function addUpstream(
  body: API.Upstream,
  options?: { [key: string]: any },
  ) {
  return request(`${API_BASE_URL}/upstream`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
    })
}

export async function removeUpstream(
  name: string,
  options?: { [key: string]: any },
  ) {
  return request(`${API_BASE_URL}/upstream/${name}`, {
    method: 'DELETE',
   ...(options || {}),
    })
}