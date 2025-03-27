import React, { useState, useEffect, useRef } from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import ProTable  from '@ant-design/pro-table';
import { Button, Modal, Form, Input, message } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { listUser, addUser, updateUser, deleteUser } from '@/services/user';
import { useIntl } from '@umijs/max';

// 定义用户类型
interface User {
  id: number;
  name: string;
  email: string;
}

// 定义服务类型
type ListUserService = () => Promise<User[]>;

const UserManagement: React.FC = () => {
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [form] = Form.useForm();
  const [editingUser, setEditingUser] = useState<number | null>(null);

  const showModal = (user: User | null = null) => {
    if (user) {
      form.setFieldsValue(user);
      setEditingUser(user.id);
    } else {
      form.resetFields();
      setEditingUser(null);
    }
    setIsModalVisible(true);
  };

  const handleOk = async () => {
    try {
      const values = await form.validateFields();
      if (editingUser) {
        await updateUser(editingUser, values);
        message.success('User updated successfully');
      } else {
        await addUser(values);
        message.success('User created successfully');
      }
      setIsModalVisible(false);
      actionRef.current?.reload();
    } catch (error) {
      console.error('Validate Failed:', error);
    }
  };

  const handleCancel = () => {
    form.resetFields();
    setIsModalVisible(false);
    setEditingUser(null);
  };

  const handleDelete = async (id: number) => {
    try {
      await deleteUser(id);
      message.success('User deleted successfully');
      // 确保 actionRef.current 存在再调用 reload 方法
      if (actionRef.current) {
        actionRef.current.reload();
      }
    } catch (error) {
      message.error('Failed to delete user');
    }
  };

  // 明确 actionRef 的类型
  const actionRef = useRef<any>();
  const onTableLoad = (action: any) => {
    actionRef.current = action;
  };

  const columns = [
    {
      title: '姓名',
      dataIndex: 'name',
    },
    {
      title: '邮件',
      dataIndex: 'email',
    },
    {
      title: '操作',
      valueType: 'option',
      render: (_, record: User) => [
        <a key="edit" onClick={() => showModal(record)}>Edit</a>,
        <a key="delete" onClick={() => handleDelete(record.id)}>Delete</a>,
      ],
    },
  ];

  const intl = useIntl();
  return (
    <PageContainer 
    content={intl.formatMessage({
      id: 'pages.user.usermanagement.informations',
      defaultMessage: '用户管理页面',
    })}
    >
      <ProTable
        headerTitle=""
        actionRef={actionRef}
        rowKey="id"
        search={false}
        toolBarRender={() => [
          <Button type="primary" onClick={() => showModal()} icon={<PlusOutlined />}>
            创建用户
          </Button>,
        ]}
        request={listUser as ListUserService}
        columns={columns}
        onLoad={onTableLoad}
      />
      <Modal
        title={editingUser ? '编辑用户' : '创建用户'}
        open={isModalVisible}
        onOk={handleOk}
        onCancel={handleCancel}
      >
        <Form form={form}>
          <Form.Item name="name" label="姓名" rules={[{ required: true, message: 'Please input your name!' }]}>
            <Input />
          </Form.Item>
          <Form.Item name="password" label="密码" rules={[{ required: true, message: 'Please input your password!' }]}>
            <Input />
          </Form.Item>
          <Form.Item name="email" label="邮箱" rules={[{ required: true, message: 'Please input your email!' }]}>
            <Input />
          </Form.Item>
        </Form>
      </Modal>
    </PageContainer>
  );
};

export default UserManagement;