import React, {Component, Fragment} from 'react';
import {DeleteProm, GetPromList, UpdateProm} from "../../api/prom";
import {Button, Form, Input, message, Modal, Pagination, Space, Table} from 'antd';

const layout = {
    labelCol: {span: 3},
    wrapperCol: {span: 15},
};

const buttonItemLayout = {
    wrapperCol: {span: 14, offset: 3},
}


class PromList extends Component {
    constructor(props) {
        super(props);
        this.state = {
            id: 0,
            index:-1,
            visible: false,
            update_visible: false,
            confirmLoading: false,
            loading: true,
            current_page: 1,
            page_size: 10,
            total: 0,
            list: [],
            column: [
                {
                    title: 'index',
                    dataIndex: 'index',
                    key: 'name',
                },
                {
                    title: 'Name',
                    dataIndex: 'name',
                    key: 'name',
                },
                {
                    title: 'Url',
                    dataIndex: 'url',
                    key: 'address',
                },
                {
                    title: 'Action',
                    key: 'action',
                    render: (text, record) => (
                        <Space size="middle">
                            <Button type="primary" className="btn-table-edit" onClick={this.onEdit.bind(this, record)}>
                                编辑
                            </Button>
                            <Button className="btn-table-edit" danger onClick={this.showModel.bind(this, record)}>
                                删除
                            </Button>
                        </Space>
                    ),
                },
            ]
        }
    }

    componentDidMount() {
        this.getData();
    }

    getData = () => {
        this.setState({
            loading: true,
        }, () => {
            const req = {
                page: this.state.current_page,
                page_size: this.state.page_size,
            }
            GetPromList(req).then(res => {
                if (res.code === 0) {
                    if (res.data.list) {
                        res.data.list.map((item, index) => {
                            return item.index = index + 1;
                        })
                        this.setState({
                            list: res.data.list,
                            total: res.data.pagination.total,
                            current_page: res.data.pagination.current_page,
                            page_size: res.data.pagination.page_size,
                        })
                    } else {
                        this.setState({
                            list: [],
                        })
                    }
                } else {
                    message.destroy();
                    message.error(res.msg);
                }
            }).catch(function (err) {
                message.destroy();
                message.error(err);
            })
            this.setState({
                loading: false,
            })
        })
    }

    onEdit({index,id, name, url}) {
        this.setState({
            update_visible: true,
            id: id,
            index:index,
        }, () => {
            this.refs.form.setFieldsValue({
                name: name,
                url: url,
            })
        })
    }

    showModel({id}) {
        this.setState({
            id: id,
            visible: true,
        })
    }

    onFinish = ({name, url}) => {
        this.setState({}, () => {
            const param = {
                id: this.state.id,
                name: name,
                url: url,
            }
            UpdateProm(param).then(res => {
                if (res.code === 0) {
                    message.destroy();
                    message.info("更新成功");
                    this.setState({
                        update_visible:false,
                    })
                    this.getData();
                }else {
                    message.destroy();
                    message.error("更新失败:"+res.msg);
                }

            }).catch(function (err) {
                message.destroy();
                message.error("更新失败:" + err);
            })
        })
    }
    onChangePagination = (page, pageSize) => {
        this.setState({
            page: page,
            page_size: pageSize,
        }, () => {
            this.getData();
        })
    }
    handleOk = () => {
        this.setState({
            confirmLoading: true,
        }, () => {
            const param = {
                id: this.state.id,
            }
            DeleteProm(param).then((res) => {
                console.log(res);
                if (res.code === 0) {
                    this.getData();
                } else {
                    message.destroy();
                    message.error(res.msg);
                }
            }).catch(function (err) {
                message.destroy();
                message.error(err);
            })
            this.setState({
                confirmLoading: false,
                visible: false,
            })
        })
    }
    handleCancel = () => {
        this.setState({
            visible: false,
        })
    }
    handleUpdateOK = () => {
        this.setState({
            update_visible: false
        })
    }
    handleUpdateCancel = () => {
        this.setState({
            update_visible: false
        })
    }

    render() {
        return (
            <Fragment>
                <Table
                    scroll={{x: 1000, y: 400}}
                    loading={this.state.loading}
                    columns={this.state.column}
                    dataSource={this.state.list}
                    rowKey={record => record.index}
                    pagination={false}
                />
                <br></br>

                <Pagination className="table-pageination"
                            showQuickJumper
                            current={this.state.current_page}
                            showSizeChanger={true}
                            total={this.state.total}
                            showTotal={(total) => `共${total}条 `}
                            onChange={this.onChangePagination}
                />

                <Modal
                    visible={this.state.update_visible}
                    onOk={this.handleUpdateOK}
                    onCancel={this.handleUpdateCancel}
                >
                    <Form ref="form"
                          {...layout}
                          onFinish={this.onFinish}
                    >
                        <Form.Item label="name" name="name"
                                   rules={[{required: true, message: '请输入数据源名称'}]}
                        >
                            <Input/>
                        </Form.Item>
                        <Form.Item label="url" name="url"
                                   rules={[{required: true, message: '请输入数据源地址'}]}
                        >
                            <Input/>
                        </Form.Item>
                        <Form.Item {...buttonItemLayout}>
                            <Button type="primary" htmlType="submit" loading={this.state.submit_loading}>
                                提交
                            </Button>
                        </Form.Item>
                    </Form>
                </Modal>

                <Modal
                    visible={this.state.visible}
                    title="警告"
                    onOk={this.handleOk}
                    confirmLoading={this.state.confirmLoading}
                    onCancel={this.handleCancel}
                >
                    <p style={{color: 'red'}}>是否确认删除该条数据源信息</p>
                </Modal>
            </Fragment>
        );
    }
}

export default PromList;