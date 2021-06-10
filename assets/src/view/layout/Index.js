import React, { Component } from 'react';
import './index.scss';
import { Layout,Button } from 'antd';
import Aside from '../../components/Aside/index';
import ContainerMain from '../../components/Container/index';
import {MenuFoldOutlined} from '@ant-design/icons';
const { Header, Content, Sider } = Layout;
class Index extends Component {
    constructor(props) {
        super(props);
        this.state = {
            collapsed: false,
        }
    }
    toggle = ()=> {
        this.setState({
            collapsed:!this.state.collapsed,
        });
    }
    Logout = ()=>{
        this.props.history.push('/antd/login');
    }
    render() {
        return (
            <Layout className="layout-warp">
                <Sider width="250px" collapsed={this.state.collapsed} theme="dark">
                    <div style={{height:"75px"}}><span>LOGO</span></div>
                    <Aside ></Aside>
                </Sider>
                <Layout>
                    <Header className="layout-header">
                        <span onClick={this.toggle}><MenuFoldOutlined /></span>
                        <h1>pormetheus告警管理页面</h1>
                        <Button className="layout-header-button" onClick={this.Logout}>退出</Button>
                    </Header>
                    <Content className="layout-content">
                        <ContainerMain></ContainerMain>
                    </Content>
                </Layout>
            </Layout>
        )
    }
}

export default Index;