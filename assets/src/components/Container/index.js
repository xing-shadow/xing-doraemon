import React, { Component } from 'react';
import {Switch} from 'react-router-dom';
import PrivateRoute from '@/components/private_route/Index';
const files = require.context("../../view/",true,/\.js$/)
const compoents = [];
files.keys().map(item => {
    if (item.includes("./layout") || item.includes("./login")) {
        return false;
    }
    let jsonObj = {};
    let splitFileName = item.split('.');
    let path = `/antd/dist${splitFileName[1].toLowerCase()}`
    let component = files(item).default;
    jsonObj.path = path;
    jsonObj.component = component;
    compoents.push(jsonObj);
    return false;
})

class ContaninerMain extends Component {
    constructor(props) {
        super(props);
        this.state = { 
            showModal: false,
         }
    }
    render() { 
        return (
            <Switch>
                {
                    compoents.map(item=> {
                        return <PrivateRoute key={item.path} exact path={item.path} component={item.component}></PrivateRoute>
                    })
                }
            </Switch>
        );
    }
}
 
export default ContaninerMain;