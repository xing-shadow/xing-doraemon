import React, { Component, Fragment } from 'react';
import { Button, Modal, Form, Input, TimePicker, InputNumber, message, Space, Table, Pagination ,Popconfirm} from 'antd';
import moment from 'moment';
import { AddPlan, GetPlan,DeletePlan } from '@/api/index';
const { RangePicker } = TimePicker;
const layout = {
    labelCol: { span: 6 },
    wrapperCol: { offset: 1 },
};

class PlanList extends Component {
    constructor(props) {
        super(props);
        this.state = {
            visible: false,
            confirmLoading: false,
            current_page: 1,
            page_size: 10,
            loading: true,
            plan_list: [],
            total: 0,
            column: [
                {
                    title: '序号',
                    dataIndex: 'index',
                    key: 'index',
                },
                {
                    title: '名称',
                    dataIndex: 'name',
                    key: 'index',
                },
                {
                    title: '开始时间',
                    dataIndex: 'start_time',
                    key: 'index',
                },
                {
                    title: '结束时间',
                    dataIndex: 'end_time',
                    key: 'index',
                },
                {
                    title: 'Fliter',
                    dataIndex: 'expression',
                    key: 'index',
                },
                {
                    title: '发送周期',
                    dataIndex: 'period',
                    key: 'index',
                },
                {
                    title: 'Action',
                    key: 'action',
                    render: (text, record) => (
                        <Space size="middle">
                            <Popconfirm
                                title="Are you sure delete this record?"
                                onConfirm={this.Deleteconfirm.bind(this,record)}
                                okText="Yes"
                                cancelText="No"
                            >
                                <Button className="btn-table-edit" danger>
                                删除
                            </Button>
                            </Popconfirm>
                        </Space>
                    ),
                },
            ]
        }
    }
    componentDidMount() {
        this.getData();
    }
    Deleteconfirm({id}) {
        this.DelePlanItem(id);
    }
    DelePlanItem = (id)=> {
        const param = {
            id:id,
        }
        DeletePlan(param).then(res => {
            if (res.code === 0) {
                message.destroy();
                message.success("删除成功");
                this.getData();
            }else{
                console.log(res);
                message.destroy();
                message.success("删除失败");
            }
        })
    }
    getData = () => {
        this.setState({
            loading: true,
        }, () => {
            const req = {
                page: this.state.current_page,
                page_size: this.state.page_size,
            }
            GetPlan(req).then(res => {
                if (res.code === 0) {
                    if (res.data.plan_list) {
                        res.data.plan_list.map((item, index) => {
                            return item.index = index + 1;
                        })
                        this.setState({
                            plan_list: res.data.plan_list,
                        })
                    }else{
                        this.setState({
                            plan_list: [],
                        })
                    }
                    this.setState({
                        total: res.data.total,
                    })
                }
            })
            this.setState({
                loading: false,
            })
        })
    }
    onShowModal = () => {
        this.setState({
            visible: true,
        })
    }
    AddhandleCancel = () => {
        this.setState({
            visible: false,
        })
    }
    AddhandleOk = () => {
        this.AddPlanItem();
    }
    AddPlanItem = () => {
        const value = this.refs.form.getFieldsValue();
        if (!value.time) {
            message.destroy();
            message.error('请选择时间范围');
            return false;
        }
        if (!value.name) {
            message.destroy();
            message.error('请输入告警计划名称')
            return false;
        }
        const param = {
            start_time: moment(value.time[0]).format("hh:mm:ss"),
            end_time: moment(value.time[1]).format("hh:mm:ss"),
            name: value.name,
            period: value.period,
            expression: value.expression,
        }
        this.setState({
            confirmLoading: true,
        }, () => {
            AddPlan(param).then(res => {
                if (res.code === 0) {
                    message.destroy();
                    message.success('添加成功');
                    this.setState({
                        visible: false,
                        confirmLoading: false,
                    })
                    this.refs.form.resetFields();
                    this.getData();
                } else {
                    message.destroy();
                    message.error('添加失败');
                    this.setState({
                        confirmLoading: false,
                    })
                }
            })
        })
    }
    onChangePagination = (page, pageSize) => {
        this.setState({
            current_page: page,
            page_size: pageSize,
        }, () => {
            this.getData();
        })
    }
    render() {
        return (
            <Fragment>
                <Button type="primary" onClick={this.onShowModal}>添加</Button>
                <br></br>
                <Table
                    scroll={{ x: 1000, y: 400 }}
                    loading={this.state.loading}
                    columns={this.state.column}
                    dataSource={this.state.plan_list}
                    rowKey={record => record.index}
                    pagination={false}
                />
                <Pagination className="table-pageination"
                    showQuickJumper
                    current={this.state.current_page}
                    showSizeChanger={true}
                    total={this.state.total}
                    showTotal={(total) => `共${total}条 `}
                    onChange={this.onChangePagination}
                />
                <Modal
                    visible={this.state.visible}
                    title="添加报警计划"
                    centered={true}
                    okText="确定"
                    cancelText="取消"
                    onCancel={this.AddhandleCancel}
                    onOk={this.AddhandleOk}
                >
                    <Form ref="form"
                    >
                        <Form.Item label="计划名称" name="name" {...layout}>
                            <Input
                                placeholder="Basic usage"
                                style={{ width: "300px" }}
                            ></Input>
                        </Form.Item>
                        <Form.Item label="报警时间段" name="time" {...layout}>
                            <RangePicker
                                style={{ width: "300px" }}
                            ></RangePicker>
                        </Form.Item>
                        <Form.Item label="Fliter" name="fliter" {...layout} initialValue="">
                            <Input
                                placeholder="Basic usage"
                                style={{ width: "300px" }}
                            ></Input>
                        </Form.Item>
                        <Form.Item label="报警周期" name="period" {...layout}
                            initialValue={1}
                        >
                            <InputNumber style={{ width: "300px" }}
                                min={1}
                                defaultValue={1}
                                formatter={value => `${value}分钟`}
                                max={60}
                            ></InputNumber>
                        </Form.Item>
                    </Form>
                </Modal>
            </Fragment>
        );
    }
}

export default PlanList;