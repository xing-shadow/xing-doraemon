import React, { Component, Fragment } from 'react';
import { Modal, Form, Button, Input,Select,message,Space ,Popconfirm,Table,InputNumber,Pagination} from 'antd';
import {GetPlanAllName} from '../../api/plan';
import {GetPromAllName} from "../../api/prom"
import {GetRule,AddRule,GetRuleList,EditRule,DeleteRule} from  '../../api/rule';
const { Option } = Select;
const layout = {
    labelCol: { span: 6 },
    wrapperCol: { offset: 1 },
};

class RuleList extends Component {
    constructor(props) {
        super(props);
        this.state = {
            id:0,
            loading:false,
            visible: false,
            current_page: 1,
            page_size: 10,
            rule_list:[],
            prom_name_list:[],
            plan_name_list:[],
            total:0,
            column: [
                {
                    title: '序号',
                    dataIndex: 'index',
                    key: 'index',
                },
                {
                    title: '指标',
                    dataIndex: 'expr',
                    key: 'index',
                },
                {
                    title: '持续时间',
                    dataIndex: 'for',
                    key: 'index',
                },
                {
                    title: 'op',
                    dataIndex: 'op',
                    key: 'index',
                },
                {
                    title: 'value',
                    dataIndex: 'value',
                    key: 'index',
                },
                {
                    title: '简介',
                    dataIndex: 'summary',
                    key: 'index',
                },
                {
                    title: '详细信息',
                    dataIndex: 'description',
                    key: 'index',
                },
                {
                    title: 'Action',
                    key: 'action',
                    render: (text, record) => (
                        <Space size="middle">
                            <Button type="primary" className="btn-table-edit" onClick={this.OnClickEditRule.bind(this,record)}>
                                编辑
                            </Button>
                            <Popconfirm
                                title="Are you sure delete this record?"
                                okText="Yes"
                                cancelText="No"
                                onConfirm={this.Deleteconfirm.bind(this,record)}
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
    componentDidMount(){
        this.GetPromNameList();
        this.GetPlanNameList();
        this.GetRuleList();
    }
    GetRuleList = () => {
        this.setState({
            loading:true,
        },()=> {
            const req = {
                page: this.state.current_page,
                page_size: this.state.page_size,
            }
            GetRuleList(req).then(res => {
                if (res.code === 0) {
                    if (res.data.rules) {
                        res.data.rules.map((item,index) => {
                            return item.index = index+1
                        })
                        this.setState({
                            rule_list:res.data.rules,
                        })
                    }else{
                        this.setState({
                            rule_list:[],
                        })
                    }
                    this.setState({
                        total:res.data.total,
                    })
                }else{
                    message.destroy();
                    message.error(res.msg);
                }
            })
            this.setState({
                loading:false,
            })
        })
    }
    OnClickEditRule(record) {
        this.setState({
            visible:true,
            id: record.id
        },() => {
            const req = {
                id:this.state.id
            }
            GetRule(req).then(res => {
                if (res.code === 0) {
                    this.refs.form.setFieldsValue({
                        metrics:res.data.expr,
                        for:res.data.for,
                        op:res.data.op,
                        plan:res.data.plan_name,
                        prom:res.data.prom_name,
                        value:res.data.value,
                        summary:res.data.summary,
                        describe:res.data.description
                    })
                }else{
                    this.setState({
                        visible:false,
                    })
                    message.destroy();
                    message.error("请求记录失败");
                }
            })
        })
    }
    Deleteconfirm({id}) {
        this.DelePlanItem(id);
    }
    DelePlanItem = (id)=> {
        const param = {
            id:id,
        }
        DeleteRule(param).then(res => {
            if (res.code === 0) {
                message.destroy();
                message.success("删除成功");
                this.GetRuleLsit();
            }else{
                console.log(res);
                message.destroy();
                message.error(res.msg);
            }
        })
    }
    GetPromNameList = ()=> {
        GetPromAllName().then(res => {
            if (res.code === 0) {
                this.setState({
                    prom_name_list:res.data,
                })
            }else{
                message.error(res.msg);
            }
        })
    }
    GetPlanNameList = ()=> {
        GetPlanAllName().then(res => {
            if (res.code === 0) {
                this.setState({
                    plan_name_list:res.data,
                })
            }else{
                message.error(res.msg);
            }
        })
    }
    AddhandleCancel = () => {
        this.setState({
            visible: false,
        })
    }
    AddhandleOk = () => {
        this.state.id ? this.EditRuleItem() : this.AddRuleItem();
    }
    EditRuleItem= ()=> {
        const data = this.refs.form.getFieldsValue();
        if (!data.for) {
            message.error("请输入持续时间");
            return false;
        }
        if (!data.metrics) {
            message.error("请输入监控指标");
            return false;
        }
        if (!data.op) {
            message.error("请输入操作类型");
            return false;
        }
        if (!data.plan) {
            message.error("请选择告警计划");
            return false;
        }
        if (!data.prom) {
            message.error("请选择数据源");
            return false;
        }
        if (!data.value) {
            message.error("请输入阈值");
            return false;
        }
        if (!data.summary) {
            message.error("请输入标题");
            return false;
        }
        const req = {
            id:this.state.id,
            expr:data.metrics,
            value:data.value,
            op:data.op,
            for:data.for,
            summary:data.summary,
            description:data.describe,
            plan_name:data.plan,
            prom_name:data.prom,
        }
        EditRule(req).then(res => {
            if (res.code === 0) {
                message.destroy();
                message.success("修改成功")
                this.setState({
                    visible:false,
                })
                this.refs.form.resetFields();
                this.GetRuleLsit();
            }else{
                message.destroy();
                message.error(res.msg); 
            }
        })
    }
    AddRuleItem = () => {
        const data = this.refs.form.getFieldsValue();
        if (!data.for) {
            message.error("请输入持续时间");
            return false;
        }
        if (!data.metrics) {
            message.error("请输入监控指标");
            return false;
        }
        if (!data.op) {
            message.error("请输入操作类型");
            return false;
        }
        if (!data.plan) {
            message.error("请选择告警计划");
            return false;
        }
        if (!data.prom) {
            message.error("请选择数据源");
            return false;
        }
        if (!data.value) {
            message.error("请输入阈值");
            return false;
        }
        if (!data.summary) {
            message.error("请输入标题");
            return false;
        }
        const req = {
            expr:data.metrics,
            value:data.value,
            op:data.op,
            for:data.for,
            summary:data.summary,
            description:data.describe,
            plan_name:data.plan,
            prom_name:data.prom,
        }
        AddRule(req).then(res => {
            if (res.code ===0 ) {
                message.success("添加成功")
                this.refs.form.resetFields();
                this.setState({
                    visible:false,
                })
                this.GetRuleLsit();
            }else{
                console.log(res);
                message.error(res.msg);
            }
        })
    } 
    onShowModal = () => {
        this.setState({
            id:0,
            visible: true,
        })
    }
    onChangePagination = (page, pageSize) => {
        this.setState({
            current_page: page,
            page_size: pageSize,
        }, () => {
            this.GetRuleLsit();
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
                    dataSource={this.state.rule_list}
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
                        <Form.Item label="监控指标" name="metrics" {...layout}
                            rules={[{ required: true }]}
                        >
                            <Input style={{ width: "300px" }}></Input>
                        </Form.Item>
                        <Form.Item label="告警阈值" {...layout}>
                            <Input.Group compact style={{ width: '300px%' }}>
                                <Form.Item  noStyle name="op" initialValue=""
                                rules={[{ required: true }]}
                                >
                                    <Select defaultValue="==" style={{ width: '20%' }}>
                                        <Option value="==" >==</Option>
                                        <Option value="!=" >!=</Option>
                                        <Option value=">=" > &gt;=</Option>
                                        <Option value="<=" > &lt;=</Option>
                                    </Select>
                                </Form.Item>
                                <Form.Item  noStyle name="value" initialValue=""
                                rules={[{ required: true }]}
                                >
                                    <Input style={{ width: '70%' }} />
                                </Form.Item>
                            </Input.Group>
                        </Form.Item>
                        <Form.Item label="持续时间" name="for" {...layout} initialValue=""
                            rules={[{ required: true }]}
                        >
                            <InputNumber style={{ width: "300px" }}
                                min={1}
                                defaultValue={1}
                                formatter={value => `${value}秒`}
                            ></InputNumber>
                        </Form.Item>
                        <Form.Item label="标题" name="summary" {...layout} initialValue=""
                        >
                            <Input style={{ width: "300px" }}></Input>
                        </Form.Item>
                           
                        <Form.Item label="描述" name="describe" {...layout} initialValue=""
                        >
                             <Input style={{ width: "300px" }}></Input>
                        </Form.Item>
                            
                        <Form.Item label="告警计划" name="plan" {...layout} initialValue=""
                            rules={[{ required: true }]}
                        >
                            <Select style={{ width: '300px' }}>
                                {
                                    this.state.plan_name_list ? this.state.plan_name_list.map((item,index) => {
                                        return <Option value={item} key={index}>{item}</Option>
                                    }) : () => {return false}
                                }
                            </Select>
                        </Form.Item>
                        <Form.Item label="数据源" name="prom" {...layout} initialValue=""
                            rules={[{ required: true }]}
                        >
                            <Select style={{ width: '300px' }}>
                            {
                                this.state.prom_name_list ? this.state.prom_name_list.map((item,index) => {
                                    return <Option value={item} key={index}>{item}</Option>
                                }) : () => {return false}
                            }
                            </Select>
                        </Form.Item>
                    </Form>
                </Modal>
            </Fragment>
        );
    }
}

export default RuleList;