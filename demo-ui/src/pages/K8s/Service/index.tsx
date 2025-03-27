import React, { useState, useEffect } from 'react';
import { Button, Modal, Form, Input, message, Select, Space } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import ProTable  from '@ant-design/pro-table';
import { SearchOutlined, PlusOutlined } from '@ant-design/icons';
import { listService, addService, removeService } from '@/services/k8s/service';

// 获取service数据请求
const fetchServiceData = async (params: {
  current?: number;
  pageSize?: number;
  name?: string;
  namespace?: string;
}) => {
  try {
    const response = await listService(params);
    return {
      data: response.data,
      total: response.total,
      success: response.success,
    };
  } catch (error) {
    console.error('获取service数据失败:', error);
    return {
      data: [],
      total: 0,
      success: false,
    };
  }
};

// 创建service请求
const createService = async (values: API.Service) => {
  try {
    await addService(values as API.Service);
    return Promise.resolve();
  } catch (error) {
    message.error(`创建service失败: ${error}`);
    return Promise.reject(error);
  }
};

// 删除service请求
const deleteService = async (name: string) => {
  try {
    await removeService(name);
    return Promise.resolve();
  } catch (error) {
    message.error(`删除上游失败: ${error}`);
    return Promise.reject(error);
  }
};

const Service: React.FC = () => {
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [form] = Form.useForm();

  const showModal = () => {
    setIsModalVisible(true);
  };

  const handleOk = async () => {
    try {
      const values = await form.validateFields();
      await createService(values);
      form.resetFields();
      setIsModalVisible(false);
      // 重新加载表格数据
      actionRef.current?.reload();
    } catch (error) {
      message.error(`创建失败: ${error}`);
    }
  };

  const handleCancel = () => {
    form.resetFields();
    setIsModalVisible(false);
  };

  const handleDelete = async (name: string) => {
    try {
      await deleteService(name);
      // 重新加载表格数据
      actionRef.current?.reload();
    } catch (error) {
      message.error(`删除失败:${error}`);
    }
  };

  const actionRef = React.useRef<any>();

  const columns = [
    {
      title: '名称',
      dataIndex: 'name',
    },
    {
      title: '类型',
      dataIndex: 'type',
    },
    {
      title: '操作',
      valueType: 'option',
      render: (_, record: API.Upstream) => (
        <Space>
          <a>查看</a>
          <a>配置</a>
          <a onClick={() => handleDelete(record.name)}>删除</a>
        </Space>
      ),
    },
  ];

  const searchConfig = {
    // 修复参数解构方式
    collapsed: false,
    optionRender: (searchConfig: any, { form }) => [
      <Button
        key="search"
        type="primary"
        icon={<SearchOutlined />}
        onClick={() => form.submit()}
      >
        {searchConfig.searchText}
      </Button>,
      <Button
        key="reset"
        onClick={() => {
          form.resetFields();
          form.submit();
        }}
        style={{ marginLeft: 8 }}
      >
        {searchConfig.resetText}
      </Button>,
    ],
    // 添加查询字段配置
    fields: [
      {
        label: '名称',
        name: 'name',
        component: <Input placeholder="请输入" />,
      },
      {
        label: '命名空间',
        name: 'type',
        component: <Input placeholder="请输入" />,
      },
    ],
    // 添加折叠按钮配置
    collapseRender: (collapsed: boolean) => (
      <span>{collapsed ? '展开' : '折叠'}</span>
    ),
  };

  return (
    <PageContainer>
      <ProTable
        headerTitle=""
        actionRef={actionRef}
        rowKey="id"
        search={searchConfig}
        request={fetchServiceData}
        columns={columns}
        pagination={{
          pageSize: 10,
          showSizeChanger: true,
          pageSizeOptions: ['10', '20', '50', '100'],
          showQuickJumper: true,
          total: 0,
        }}
        params={{
          name: '',
          type: '',
        }}
        toolBarRender={() => [
        <Button type="primary" onClick={showModal} icon={<PlusOutlined />}>
          创建服务
        </Button>,
        ]}
      />
      <Modal
        title="创建服务"
        open={isModalVisible}
        onOk={handleOk}
        onCancel={handleCancel}
      >
        <Form form={form}>
          <Form.Item
            name="name"
            label="名称"
            rules={[{ required: true, message: '请输入名称!' }]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="description"
            label="描述"
            rules={[{ required: true, message: '请输入描述信息!' }]}
          >
            <Input.TextArea rows={4} />
          </Form.Item>
        </Form>
      </Modal>
    </PageContainer>
  );
};

export default Service;