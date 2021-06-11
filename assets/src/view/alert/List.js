import React, { Component ,Fragment} from 'react';
import {message, Table,Button,Pagination} from 'antd';
import {GetAlertList,ConfirmAlert} from '../../api/alert';
class AlertList extends Component {
    constructor(props) {
        super(props);
        this.state = {
            current_page:1,
            page_size:10,
            loading:false,
            confirm_loading:false,
            alert_list:[],
            selectedRowKeys:[],
            total:0,
            column: [
                {
                    title: '序号',
                    dataIndex: 'index',
                    key: 'id',
                },
                {
                    title: 'summary',
                    dataIndex: 'summary',
                    key: 'id',
                },
                {
                    title: 'instance',
                    dataIndex: 'instance',
                    key: 'id',
                },
                {
                    title: 'value',
                    dataIndex: 'value',
                    key: 'id',
                },
                {
                    title: 'fired_at',
                    dataIndex: 'fired_at',
                    key: 'id',
                },
                {
                    title: 'count',
                    dataIndex: 'count',
                    key: 'id',
                },
                {
                    title: 'labels',
                    dataIndex: 'labels',
                    key: 'id',
                },
            ]
        }
    }
    componentDidMount() {
       this.getData();
    }
    getData= ()=> {
        this.setState({
            loading:true,
        },()=> {
            const req = {
                page:this.state.current_page ? this.state.current_page: 1,
                page_size:this.state.page_size,
            }
            GetAlertList(req).then(res => {
                if (res.code === 0) {
                    if (res.data) {
                        if (res.data.alerts) {
                            res.data.alerts.map((item,index) => {
                                return item.index = index+1;
                            })
                        }
                        this.setState({
                            alert_list:res.data.alerts,
                            total:res.data.total
                        })
                    }else{
                        this.setState({
                            alert_list:[],
                            total:0,
                        }) 
                    }
                }else{
                    message.destroy();
                    message.error(res.msg)
                }
            })
            this.setState({
                loading:false
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
    onSelectChange = (selectedRowKeys) => {
        this.setState({
            selectedRowKeys:selectedRowKeys,
        })
    }
    onConfirmAlert = ()=> {
        if (this.state.selectedRowKeys.length <= 0 ) {
            return false;
        }
        const req = {
            alert_list:this.state.selectedRowKeys,
        }
        this.setState({
            confirm_loading:true,
        },()=> {
            ConfirmAlert(req).then(res => {
                if (res.code === 0) {
                    message.destroy();
                    message.success("确认告警成功");
                    this.getData();
                }else{
                    message.destroy();
                    message.error(res.msg);
                }
                this.setState({
                    confirm_loading:false,
                })
            })
        })
    }
    render() { 
        const { selectedRowKeys } = this.state;
        const rowSelection = {
            selectedRowKeys,
            onChange: this.onSelectChange,
          };
        return (
        <Fragment>
            <Button type="primary" onClick={this.onConfirmAlert} loading={this.confirm_loading}>确认告警</Button>
            <br></br>
            <Table
                scroll={{ x: 1000, y: 400 }}
                loading={this.state.loading}
                columns={this.state.column}
                dataSource={this.state.alert_list}
                rowKey={record => record.index}
                rowSelection= {
                    rowSelection
                }
                pagination={false}
            >
            </Table>
            <Pagination className="table-pageination"
                showQuickJumper
                current={this.state.current_page}
                showSizeChanger={true}
                total={this.state.total}
                showTotal={(total) => `共${total}条 `}
                onChange={this.onChangePagination}
            />
        </Fragment>
        );
    }
}
 
export default AlertList;