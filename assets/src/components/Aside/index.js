import React, { Component, Fragment } from 'react';
import { Menu } from 'antd';
import Router from '../../router/index';
import { Link } from 'react-router-dom';
import { UserOutlined,AppstoreOutlined } from '@ant-design/icons';

const { SubMenu } = Menu;

class Aside extends Component {
    constructor(props) {
        super(props);
        this.state = {

        }
    }

    renderMenuItem = (item) => {
        return (
            <Menu.Item key={item.key}>
                <UserOutlined></UserOutlined>
                <Link to={item.key}>{item.title}</Link>
            </Menu.Item>
        )
    }

    renderSubMenu = (items) => {
        return (
            <SubMenu key={items.key} title={items.title} icon={<AppstoreOutlined></AppstoreOutlined>}>
            {
                items.subs && items.subs.map(item => {
                    return (item.subs && item.subs.length > 0) ? this.renderSubMenu(item) : this.renderMenuItem(item)
                })
            }
            </SubMenu>
        )
    }

    render() {
        return (
            <Fragment>
                <Menu
                    defaultSelectedKeys={['1']}
                    defaultOpenKeys={['sub1']}
                    mode="inline"
                    theme="dark"
                >
                    {
                        Router && Router.map(item => {
                            return (item.subs && item.subs.length > 0) ? this.renderSubMenu(item) : this.renderMenuItem(item)
                        })
                    }
                </Menu>
            </Fragment>
        );
    }
}

export default Aside;