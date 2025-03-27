import { request } from "@umijs/max";

export async function listDeployment(
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
  return request<API.DeploymentList>(`${API_BASE_URL}/deployment`, {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
    })
  }

  export async function addDeployment(
    body: API.Deployment,
    options?: { [key: string]: any },
    ) {
    return request<API.Deployment>(`${API_BASE_URL}/deployment`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: body,
      ...(options || {}),
    })
  }