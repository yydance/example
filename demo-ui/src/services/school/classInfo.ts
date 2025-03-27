import { request } from '@umijs/max';

export async function listClass(
  params?: API.classQuery,
  options?: { [key: string]: any },
) {
  return request<API.classList>('/eeos_admin/action/ops/course_single.php', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}