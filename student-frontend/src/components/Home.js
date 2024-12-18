import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import { FaUserGraduate, FaUserPlus, FaList, FaChartBar, FaGraduationCap } from 'react-icons/fa';
import './Home.css';

function Home() {
    useEffect(() => {
        /* 初始化粒子动画 */
        const script = document.createElement('script');
        script.src = 'https://cdn.jsdelivr.net/npm/particles.js@2.0.0/particles.min.js';
        script.onload = () => {
            window.particlesJS('particles-js', {
                particles: {
                    number: { value: 80 },
                    color: { value: '#ffffff' },
                    shape: { type: 'circle' },
                    opacity: { value: 0.5 },
                    size: { value: 3 },
                    line_linked: {
                        enable: true,
                        distance: 150,
                        color: '#ffffff',
                        opacity: 0.4,
                        width: 1
                    },
                    move: { enable: true, speed: 2 }
                }
            });
        };
        document.body.appendChild(script);
    }, []);

    return (
        <div className="home-container">
            <div className="hero-section">
                {/* 添加粒子动画容器 */}
                <div id="particles-js"></div>
                <div className="hero-content">
                    <FaGraduationCap className="hero-icon" />
                    <h1>智慧校园管理系统</h1>
                    <p className="hero-subtitle">简单 · 高效 · 智能</p>
                    <div className="hero-stats">
                        <div className="stat-item">
                            <span className="stat-number">1000+</span>
                            <span className="stat-label">学生数据</span>
                        </div>
                        <div className="stat-item">
                            <span className="stat-number">99%</span>
                            <span className="stat-label">准确率</span>
                        </div>
                        <div className="stat-item">
                            <span className="stat-number">24/7</span>
                            <span className="stat-label">在线服务</span>
                        </div>
                    </div>
                </div>
                <div className="hero-bottom-effect">
                    <div className="light-beam">
                        <div className="floating-particle"></div>
                    </div>
                    <div className="light-beam">
                        <div className="floating-particle"></div>
                    </div>
                    <div className="light-beam">
                        <div className="floating-particle"></div>
                    </div>
                    <div className="light-beam">
                        <div className="floating-particle"></div>
                    </div>
                    <div className="light-beam">
                        <div className="floating-particle"></div>
                    </div>
                </div>
            </div>

            <div className="features-section">
                <h2 className="section-title">核心功能</h2>
                <div className="features-grid">
                    <Link to="/student-list" className="feature-card">
                        <FaList className="feature-icon" />
                        <h3>学生列表</h3>
                        <p>查看和管理所有学生信息</p>
                    </Link>

                    <Link to="/add-student" className="feature-card">
                        <FaUserPlus className="feature-icon" />
                        <h3>添加学生</h3>
                        <p>录入新学生信息</p>
                    </Link>

                    <Link to="/system-overview" className="feature-card">
                        <FaUserGraduate className="feature-icon" />
                        <h3>系统概况</h3>
                        <p>查看学生数据统计信息</p>
                    </Link>
                </div>
            </div>

            <div className="workflow-section">
                <h2 className="section-title">使用流程</h2>
                <div className="workflow-steps">
                    <div className="step">
                        <div className="step-number">1</div>
                        <div className="step-content">
                            <h3>录入信息</h3>
                            <p>快速添加学生基本信息</p>
                        </div>
                    </div>
                    <div className="step">
                        <div className="step-number">2</div>
                        <div className="step-content">
                            <h3>数据管理</h3>
                            <p>便捷的信息维护更新</p>
                        </div>
                    </div>
                    <div className="step">
                        <div className="step-number">3</div>
                        <div className="step-content">
                            <h3>统计分析</h3>
                            <p>智能数据分析报告</p>
                        </div>
                    </div>
                </div>
            </div>

            <div className="info-section">
                <div className="info-card">
                    <h4>便捷管理</h4>
                    <p>快速录入和更新学生信息</p>
                </div>
                <div className="info-card">
                    <h4>数据安全</h4>
                    <p>安全可靠的数据存储和处理</p>
                </div>
                <div className="info-card">
                    <h4>简单易用</h4>
                    <p>直观的用户界面设计</p>
                </div>
            </div>
        </div>
    );
}

export default Home;
