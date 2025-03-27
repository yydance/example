import { PageContainer } from '@ant-design/pro-components';
import { useIntl } from '@umijs/max';
import { DownOutlined, UpOutlined } from '@ant-design/icons';
import { Input, Card, message, Select, Button, DatePicker, Form, Row, Col } from 'antd';
import React, { useEffect, useState } from 'react';
import ProTable  from '@ant-design/pro-table';
import dayjs from 'dayjs';
import { listClass } from '@/services/school/classInfo';

const classStatus: { [key: string]: number } = {
  'pages.lesson-list.class-status-text.1': 1,
  'pages.lesson-list.class-status-text.2': 2,
  'pages.lesson-list.class-status-text.3': 3,
  'pages.lesson-list.class-status-text.4': 4
}

const fetchClassData = async (
  params?: API.classQuery ) => {
  try {
    //console.log('[DEBUG] 请求参数:', {
    //  ...params,
    //});
    const response = await listClass(params);
    return {
      data: response.data,
      success: true,
      total: response.data?.length || 0,
    };
  } catch (error) {
    const intl = useIntl();
    const lessonFailed = intl.formatMessage({ id: 'pages.schooladmin.lesson-list.getLessonFailed' });
    message.error(`${lessonFailed}: ${error}`)
    return {
      data: [],
      success: false,
      total: 0,
    };
  };
};

const Lesson: React.FC = () => {
  const intl = useIntl();
  const STORAGE_KEY = 'lesson_search_params';
  // 添加默认时间范围
  const defaultDateRange = [
    dayjs().startOf('day'),
    dayjs().endOf('day'),
  ];
  const { RangePicker } = DatePicker;
  const [form] = Form.useForm();

  const [searchParams, setSearchParams] = useState<API.classQuery>({});
  const [collapsed, setCollapsed] = useState(true);
  const [dateRange, setDateRange] = useState(defaultDateRange);

  // 初始化时读取存储的搜索参数
  useEffect(() => {
    const savedParams = localStorage.getItem(STORAGE_KEY);
    if (savedParams) {
      const params = JSON.parse(savedParams);
      if (params.time_range) {
        params.time_range = params.time_range.map((t: string) => dayjs(t));
      }
      setSearchParams(params);
      form.setFieldsValue(params);
    }
  }, []);

  const onFinish = (values: any) => {
    const { current, page_size, ...restValues } = values;
    values.time_range = dateRange;
    //console.log('表单值：',values);
    const timeRangeStr = values.time_range?.map((t: dayjs.Dayjs) => 
        t.format('YYYY-MM-DD')
      ).join(',') || '';
      localStorage.setItem(STORAGE_KEY, JSON.stringify(restValues));
      setSearchParams({
        ...restValues,
        time_str: timeRangeStr
      });
    }
  const handleDateChange = (dates: any) => {
    form.setFieldsValue({ time_range: dates });
    const currentValues = form.getFieldsValue();
    const storageValues = {
      ...currentValues,
      time_range: dates?.map((t: any) => t.toISOString()) || []
    };
    localStorage.setItem(STORAGE_KEY, JSON.stringify(storageValues));
    setDateRange(dates || defaultDateRange);
  };

  const columns = [
    {
      title: 'pages.lesson-list.class-btime',
      dataIndex: 'class_btime',
    },
    {
      title: 'pages.lesson-list.class-etime',
      dataIndex:'class_etime',
    },
    {
      title: 'pages.lesson-list.class-id',
      dataIndex: 'class_id',
    },
    {
      title: 'pages.lesson-list.class-name',
      dataIndex: 'class_name',
    },
    {
      title: 'pages.lesson-list.school-uid',
      dataIndex: 'school_uid',
    },
    {
      title: 'pages.lesson-list.school-name',
      dataIndex: 'school_name',
    },
    {
      title: 'pages.lesson-list.course-id',
      dataIndex: 'course_id',
    },
    {
      title: 'pages.lesson-list.course-name',
      dataIndex: 'course_name',
    },
    {
      title: 'pages.lesson-list.main-teacher',
      dataIndex: 'main_teacher',
    },
    {
      title: 'pages.lesson-list.assistants',
      dataIndex: 'assistants',
    },
    {
      title: 'pages.lesson-list.student-list-url',
      dataIndex: 'student_list_url',
    },
    {
      title: 'pages.lesson-list.class-status',
      dataIndex: 'class_status',
    },
    {
      title: 'pages.lesson-list.client-class-id',
      dataIndex: 'client_class_id',
    },
    {
      title: 'pages.lesson-list.seat-num',
      dataIndex: 'seat_num',
    },
    {
      title: 'pages.lesson-list.folder-name',
      dataIndex: 'folder_name',
    },
    {
      title: 'pages.lesson-list.folder-path',
      dataIndex: 'folder_path',
    },
    {
      title: 'pages.lesson-list.record-state',
      dataIndex:'record_state',
    },
    {
      title: 'pages.lesson-list.live-state',
      dataIndex:'live_state',
    },
    {
      title: 'pages.lesson-list.open-state',
      dataIndex:'open_state',
    },
    {
      title: 'pages.lesson-list.student-num',
      dataIndex:'student_num',
    },
    {
      title: 'pages.lesson-list.class-status-text',
      dataIndex:'class_status_text',
    },
    {
      title: 'pages.lesson-list.st-id',
      dataIndex:'st_id',
    },
    {
      title: 'pages.lesson-list.ass-st-id',
      dataIndex:'ass_st_id',
    },
    {
      title: 'pages.lesson-list.addtime',
      dataIndex:'addtime',
    },
    {
      title: 'pages.lesson-list.record-url',
      dataIndex: 'record_url',
    },
  ];

  const defaultColumns = {
    class_btime: {show: true},
    class_etime: {show: true},
    class_id: {show: true},
    class_name: {show: true},
    school_uid: {show: true},
    school_name: {show: true},
    course_id: {show: true},
    course_name: {show: true},
    main_teacher: {show: true},
    assistants: {show: true},
    student_list_url: {show: false},
    class_status: {show: false},
    client_class_id: {show: false},
    seat_num: {show: false},
    folder_name: {show: false},
    folder_path: {show: false},
    record_state: {show: false},
    live_state: {show: false},
    open_state: {show: false},
    student_num: {show: false},
    class_status_text: {show: false},
    st_id: {show: false},
    ass_st_id: {show: false},
    addtime: {show: false},
    record_url: {show: false},
  }

  const renderSearchBar = () => (
    < div 
    className="ant-pro-table-search" 
    style={{ 
      marginBottom: 16,
      padding: 24,
      background: '#fff',
      border: '1px solid #f0f0f0',
      borderRadius: 2
    }}
    >
      <Form form={form} onFinish={onFinish}>
        <Row gutter={16}>
        <Col span={6}>
          <Form.Item name="time_range" label={intl.formatMessage({ id: 'pages.lesson-list.search.time-range'})}>
          <RangePicker 
            defaultValue={[
              defaultDateRange[0],
              defaultDateRange[1],
            ]}
            onChange={handleDateChange}
          />
          </Form.Item>
        </Col>
        <Col span={6}>
            <Form.Item name="class_status" label={intl.formatMessage({ id: 'pages.lesson-list.class-status'})}>
              <Select placeholder={intl.formatMessage({ id: 'pages.lesson-list.class-status.placeholder'})}>
                {Object.entries(classStatus).map(([label, value]) => (
                  <Select.Option key={value} value={value}>
                    {intl.formatMessage({ id: `${label}` })}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
          </Col>
          <Col span={6}>
            <Form.Item name="class_id" label={intl.formatMessage({ id: 'pages.lesson-list.class-id'})}>
              <Input placeholder={intl.formatMessage({ id: 'pages.lesson-list.class-id.placeholder'})} allowClear />
            </Form.Item>
          </Col>
          <Col span={6}>
            <Form.Item name="course_id" label={intl.formatMessage({ id: 'pages.lesson-list.course-id'})}>
              <Input placeholder={intl.formatMessage({ id: 'pages.lesson-list.course-id.placeholder'})} allowClear />
            </Form.Item>
          </Col>
          <Col span={6}>
            <Form.Item name="school_uid" label={intl.formatMessage({ id: 'pages.lesson-list.school-uid'})}>
              <Input placeholder={intl.formatMessage({ id: 'pages.lesson-list.school-uid.placeholder'})} allowClear />
            </Form.Item>
          </Col>
          {!collapsed && (
          <>
            <Col span={6}>
              <Form.Item name="class_name" label={intl.formatMessage({ id: 'pages.lesson-list.class-name'})}>
                <Input placeholder={intl.formatMessage({ id: 'pages.lesson-list.class-name.placeholder'})} allowClear />
              </Form.Item>
            </Col>
            <Col span={6}>
              <Form.Item name="st_id" label={intl.formatMessage({ id: 'pages.lesson-list.st-id'})}>
                <Input placeholder={intl.formatMessage({ id: 'pages.lesson-list.st-id.placeholder'})} allowClear />
              </Form.Item>
            </Col>
            <Col span={6}>
              <Form.Item name="ass_st_id" label={intl.formatMessage({ id: 'pages.lesson-list.ass-st-id'})}>
                <Input placeholder={intl.formatMessage({ id: 'pages.lesson-list.ass-st-id.placeholder'})} allowClear />
              </Form.Item>
            </Col>
            <Col span={6}>
              <Form.Item name="client_class_id" label={intl.formatMessage({ id: 'pages.lesson-list.client-class-id'})}>
                <Input placeholder={intl.formatMessage({ id: 'pages.lesson-list.client-class-id.placeholder'})} allowClear />
              </Form.Item>
            </Col>
          </>
          )}
          <Col
            span={6} 
            style={{ 
              display: 'flex',
              justifyContent: 'flex-end',
              gap: 8,
              marginLeft: 'auto' // 推动到最右侧
            }}
          >
          <Button
              type="link"
              onClick={() => setCollapsed(!collapsed)}
              icon={collapsed ? <DownOutlined /> : <UpOutlined />}
              style={{ marginRight: 8 }}
            >
              {collapsed
              ? intl.formatMessage({ id: 'component.tagSelect.expand' })
              : intl.formatMessage({ id: 'component.tagSelect.collapse' })}
            </Button>
            <Button 
              type="primary"
              htmlType="submit"
              style={{ marginRight: 8 }}
            >
              {intl.formatMessage({ id: 'component.tagSearch.search' })}
            </Button>
            <Button onClick={() => {
              form.resetFields();
              setSearchParams({});
              localStorage.removeItem(STORAGE_KEY);
            }}>
              {intl.formatMessage({ id: 'component.tagSearch.reset' })}
            </Button>
          </Col>
        </Row>
      </Form>
    </div>
  );

  return (
    <PageContainer header={{ title: false }}>
      {renderSearchBar()}
      <ProTable 
        columns={columns.map(column => ({
          ...column,
          title: intl.formatMessage({ id: column.title }),
        }))}
        columnsState={{
          defaultValue: defaultColumns,
          persistenceKey: 'lesson_columns',
          persistenceType: 'localStorage',
        }}
        request={async () => {
          // 转换分页参数
          //const { current: page = 1, page_size = 20 } = params;
          return fetchClassData({
            ...searchParams
          });
        }}
        rowKey="class_id"
        pagination={{
          showQuickJumper: true,
          pageSizeOptions: [10, 20, 50, 100],
          defaultPageSize: 10,
          pageSize: 10,
          showSizeChanger: true,
        }}
        search={false}
        params={searchParams}
      />
    </PageContainer>
  );
};

export default Lesson;
