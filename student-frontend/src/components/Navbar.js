
import React from 'react';
import './Navbar.css';

function Navbar() {
    return (
        <nav className="navbar">
            <div className="navbar-logo">学生管理系统</div>
            <ul className="navbar-menu">
                <li><a href="#">主页</a></li>
                <li><a href="#student-list">学生列表</a></li>
                <li><a href="#add-student">添加学生</a></li>
            </ul>
        </nav>
    );
}

export default Navbar;