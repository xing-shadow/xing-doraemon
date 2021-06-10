import React, { Component } from 'react';
import { Form, Input, Button ,message} from 'antd';
import {UserOutlined, LockOutlined} from '@ant-design/icons';
import './index.scss';
import { LoginUser } from '../../api/index.js';
class UserLogin extends Component {
    constructor(props) {
        super(props);
        this.state = {}
    }
    onFinish = ({ username, password }) => {
        const req = {
            username: username,
            password: password,
        }
        LoginUser(req).then(res => {
            if (res.code === 0) {
                this.props.history.push("/antd/dist/prom/list");
            }else if (res.code === 302){
                this.props.history.push(res.msg);
            }else {
                message.destroy();
                message.error("登录失败");
            }
        }).catch(function (err){
            message.destroy();
            message.error("请求失败");
        });
    }
    render() {
        return (
            <div className="form-warp">
                <div>
                    <div className="form-header">
                        <h4 className="column">登录</h4>
                    </div>
                    <div className="form-content">
                        <Form
                            name="normal_login"
                            className="login-form"
                            initialValues={{ remember: true }}
                            onFinish={this.onFinish}
                        >
                            <Form.Item
                                name="username"
                                rules={
                                    [
                                        { required: true, message: 'Please input your username!' },
                                    ]
                                }
                            >
                                <Input prefix={<UserOutlined className="site-form-item-icon" />} placeholder="username" />
                            </Form.Item>
                            <Form.Item
                                name="password"
                                rules={
                                    [
                                        { required: true, message: '密码不能为空' },

                                    ]
                                }
                            >
                                <Input.Password prefix={<LockOutlined className="site-form-item-icon" />} placeholder="password" />
                            </Form.Item>
                            <Form.Item>
                                <Button type="primary" htmlType="submit" className="login-form-button" block>
                                    登录
                                </Button>
                            </Form.Item>
                        </Form>
                    </div>
                </div>
            </div>
        );
    }
}

export default UserLogin;