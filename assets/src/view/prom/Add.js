import React, { Component,Fragment } from 'react';
import {Button, Form,Input,message} from 'antd';
import {GetProm,UpdataProm,AddProm} from '@/api/index';
import PropTypes from 'prop-types';

const layout = {
    labelCol: { span: 2 },
    wrapperCol: { span: 8 },
  };
  const tailLayout = {
    wrapperCol: { offset: 2, span: 14 },
  };

class PromAdd extends Component {
    constructor(props) {
        super(props);
        this.state = {
            id:0,
        }
    }
    UNSAFE_componentWillMount () {
        if (this.props.location.state) {
            this.setState({
                id:this.props.location.state.id
            })
        }
    }
    componentDidMount() {
        if (!this.state.id) {
            return false
        }
        this.GetPromItem();
     }
    GetPromItem = ()=> {
        const param = {
            id: this.state.id,
        }
        GetProm(param).then(res => {
            if (res.code === 0) {
                this.refs.form.setFieldsValue({
                    name:res.data.name,
                    url:res.data.url,
                })
            }else{
                message.destroy();
                message.error('请求失败')
            }
        })
    }
    onFinish = ({name,url})=> {
        !this.state.id ? this.AddPormItem(name,url):this.upDataPromItem(name,url)
    }
    AddPormItem = (name,url) => {
        const param = {
            name:name,
            url:url,
        }
        AddProm(param).then(res => {
            if (res.code === 0) {
                this.refs.form.resetFields();
                message.destroy();
                message.success("添加成功")
            }else{
                message.destroy();
                message.error("添加失败")
            }
        })
    }
    upDataPromItem = (name,url)=> {
        const param = {
            id: this.state.id,
            name: name,
            url:url
        }
        UpdataProm(param).then(res => {
            if (res.code === 0) {
                this.props.history.push({
                    pathname:"/antd/dist/prom/list",
                })
                message.destroy();
                message.success("更新成功")
            }else{
                message.destroy();
                message.success("更新失败")
            }
        })
    }
    render() { 
        return ( 
            <Fragment>
                <Form ref="form"
                {...layout}
                onFinish={this.onFinish}
                >
                    <Form.Item label="name" name="name"
                    rules={[{ required: true, message: '请输入数据源名称' }]}
                    >
                    <Input></Input>
                    </Form.Item>
                    <Form.Item label="url" name="url"
                    rules={[{ required: true, message: '请输入数据源地址' }]}
                    >
                    <Input></Input>
                    </Form.Item>
                    <Form.Item {...tailLayout}>
                        <Button type="primary" htmlType="submit">
                            提交
                        </Button>
                    </Form.Item>
                </Form>
            </Fragment>
         );
    }
}
 
export default PromAdd;

PromAdd.propTypes = {
    id: PropTypes.number,
}

PromAdd.defaultProps = {
    id: 0,
}