import React from 'react';
import {Route,Redirect} from 'react-router-dom';
import {GetToken} from '@/utils/Token';

const PrivateRoute = ({component:Component,...ret}) => {
    return (<Route {...ret} render={routeProps =>(GetToken() ? <Component {...routeProps}/> : <Redirect to="/antd/login"/>)}></Route>)
} 
export default PrivateRoute;