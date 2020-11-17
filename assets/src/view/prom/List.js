import React, { Component, Fragment } from 'react';
import { GetPormList,DeleProm } from "../../api/index";
import { Table, Space,Button,Form,Input,Pagination,Modal,message  } from 'antd';

class PromList extends Component {
    constructor(props) {
        super(props);
        this.state = {
            
            search_name:'',
            id:0,
            visible: false,
            confirmLoading:false,
            current_page: 1,
            page_size:10,
            loading: true,
            PromList: [],
            total:0,
            column:[
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
                            <Button type="primary" className="btn-table-edit" onClick={this.onEdit.bind(this,record)}>
                                编辑
                            </Button>
                            <Button className="btn-table-edit" danger onClick={this.showModel.bind(this,record)}>
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
    getData = ()=> {
        
        this.setState({
            loading: true,
        },()=>{
            const req = {
                page: this.state.current_page,
                page_size: this.state.page_size,
                name: this.state.search_name,
            }
            GetPormList(req).then(res => {
                if (res.code === 0) {
                    if (res.data.prom_list) {
                        res.data.prom_list.map((item, index) => {
                            return item.index = index + 1;
                        })
                        this.setState({
                        PromList: res.data.prom_list,
                        })
                    }else{
                        this.setState({
                            PromList: [],
                            })
                    }
                    this.setState({
                        total: res.data.total,
                    })
                }
            })
            this.setState({
                loading:false,
            })
        })
    }
    showModel({id}) {
        this.setState({
            id:id,
            visible:true,
        })
    }
    onFinish = ({name}) => {
        this.setState({
            search_name:name,
        },() => {
            this.getData();
        })
    }
    onEdit (record) {
        this.props.history.push({
            pathname:"/antd/dist/prom/add",
            state:{
                id:record.id
            }
        })
    }
    onChangePagination= (page, pageSize) => {
        this.setState({
            page:page,
            page_size:pageSize,
        },() => {
            this.getData();
        })
    }
    handleOk = () => {
        this.setState({
            confirmLoading:true,
        },()=>{
            const param = {
                id:this.state.id,
            }
            DeleProm(param).then((res) => {
                console.log(res);
                if (res.code === 0) {
                    message.destroy();
                    message.success("请求成功");
                    this.getData();
                }else{
                    console.log(res);
                    message.destroy();
                    message.error("请求失败");
                }
            })
            this.setState({
                confirmLoading:false,
                visible:false,
            })
        })
    }
    handleCancel = () => {
        this.setState({
            visible:false,
        })
    }
    render() {
        return (
            <Fragment>
                <Form
                    layout="inline"
                    onFinish={this.onFinish}
                >
                    <Form.Item label="名称" name="name">
                        <Input></Input>
                    </Form.Item>
                    <Form.Item >
                        <Button type="primary" htmlType="submit">搜索</Button>
                    </Form.Item>
                </Form>
                <br></br>
                <Table
                    scroll={{ x: 1000, y: 400 }} 
                    loading={this.state.loading}
                    columns={this.state.column} 
                    dataSource={this.state.PromList} 
                    rowKey={record => record.name}
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
                    visible={this.state.visible}
                    title="警告"
                    onOk={this.handleOk}
                    confirmLoading={this.state.confirmLoading}
                    onCancel={this.handleCancel}
                    >
                    <p style={{color:'red'}}>是否确认删除该条数据源信息</p>
                </Modal>
            </Fragment>
        );
    }
}

export default PromList;