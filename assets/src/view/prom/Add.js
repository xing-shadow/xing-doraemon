import React, {Component, Fragment} from 'react';
import {Button, Form, Input, message} from 'antd';
import {AddProm} from '../../api/prom';
import PropTypes from 'prop-types';

const layout = {
    labelCol: {span: 2},
    wrapperCol: {span: 8},
};
const tailLayout = {
    wrapperCol: {offset: 2, span: 14},
};

class PromAdd extends Component {
    constructor(props) {
        super(props);
        this.state = {
            submit_loading: false,
        }
    }

    onFinish = ({name, url}) => {
        this.AddPromItem(name, url);
    }
    AddPromItem = (name, url) => {
        this.setState({
            submit_loading: true,
        }, () => {
            const param = {
                name: name,
                url: url,
            }
            AddProm(param).then(res => {
                if (res.code === 0) {
                    this.refs.form.resetFields();
                    message.destroy();
                    message.success("添加成功")
                } else {
                    message.destroy();
                    message.error("添加失败:" + res.msg);
                }
            }).catch(function (err) {
                message.destroy();
                message.error("添加失败")
            })
            this.setState({
                submit_loading: false
            })
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
                               rules={[{required: true, message: '请输入数据源名称'}]}
                    >
                        <Input></Input>
                    </Form.Item>
                    <Form.Item label="url" name="url"
                               rules={[{required: true, message: '请输入数据源地址'}]}
                    >
                        <Input></Input>
                    </Form.Item>
                    <Form.Item {...tailLayout}>
                        <Button type="primary" htmlType="submit" loading={this.state.submit_loading}>
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