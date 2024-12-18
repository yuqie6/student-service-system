import React from 'react';
import { Link } from 'react-router-dom';
import './Navbar.css';

function Navbar() {
    return (
        <nav className="navbar">
            <div className="navbar-logo">学生管理系统</div>
            <ul className="navbar-menu">
                <li><Link to="/">主页</Link></li>
                <li><Link to="/student-list">学生列表</Link></li>
                <li><Link to="/add-student">添加学生</Link></li>
                <li><Link to="/system-overview">系统概况</Link></li>
            </ul>
        </nav>
    );
}

export default Navbar;