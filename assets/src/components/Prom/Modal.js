import React, { Component, Fragment } from 'react';
import { Modal, Button,Input } from 'antd';

class Modals extends Component {
    constructor(props) {
        super(props);
        this.state = {
            loading: false,
            visible: false,
            item: props.item,
            name:'',
            url:'',
        }
    }
    UNSAFE_componentWillMount() {
        this.setState({
            visible: this.props.visible,
        });
    }
    
    handleOk = () => {
        this.props.fromSon(false);
        this.setState({
            item:{
                name:this.name,
                url:this.url
            }
        })
        this.setState({
            visible: false,
        });
    }

    handleCancel = () =>{
        this.props.fromSon(false);
        this.setState({
            visible: false,
        });
    }
    nameChange = (e) => {
        this.setState(
            {
                name: e.target.value
            }
        )
    }
    urlChange = (e) => {
        this.setState(
            {
                url: e.target.value
            }
        )
    }
    render() {
        const { visible, loading } = this.state;
        return (
            <Fragment>
                <Modal
                    visible={visible}
                    title="编辑"
                    width="420px"
                    onOk={this.handleOk}
                    onCancel={this.handleCancel}
                    footer={[
                        <Button key="submit" type="primary" loading={loading} onClick={this.handleOk}>
                            确定
                        </Button>,
                    ]}
                >
                    <Input.Group size="large">
                        <Input addonBefore="名称" defaultValue={this.state.item.name} onChange={this.nameChange}/>
                        <br></br>
                        <br></br>
                        <Input addonBefore="url" defaultValue={this.state.item.url} onChange={this.urlChange}/>
                    </Input.Group>
                </Modal>
            </Fragment>
        );
    }
}

export default Modals;