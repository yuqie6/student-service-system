import React, { useState, useEffect } from 'react';
import { studentApi } from '../api/studentApi';
import { FaUsers, FaMale, FaFemale, FaChartBar } from 'react-icons/fa';
import './SystemOverview.css';

function SystemOverview() {
    const [stats, setStats] = useState({
        totalStudents: 0,
        maleCount: 0,
        femaleCount: 0,
        averageAge: 0
    });
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchStats();
    }, []);

    const fetchStats = async () => {
        try {
            const response = await studentApi.getAllStudents();
            const students = response.data;
            
            const maleStudents = students.filter(s => s.gender === '男');
            const femaleStudents = students.filter(s => s.gender === '女');
            const totalAge = students.reduce((sum, student) => sum + student.age, 0);

            setStats({
                totalStudents: students.length,
                maleCount: maleStudents.length,
                femaleCount: femaleStudents.length,
                averageAge: students.length ? Math.round(totalAge / students.length) : 0
            });
        } catch (error) {
            console.error('获取统计数据失败:', error);
        } finally {
            setLoading(false);
        }
    };

    if (loading) {
        return <div className="loading">加载中...</div>;
    }

    return (
        <div className="system-overview">
            <h2>系统概况</h2>
            <div className="stats-grid">
                <div className="stat-card">
                    <FaUsers className="stat-icon" />
                    <div className="stat-info">
                        <h3>总学生数</h3>
                        <p>{stats.totalStudents}</p>
                    </div>
                </div>
                <div className="stat-card">
                    <FaMale className="stat-icon" />
                    <div className="stat-info">
                        <h3>男生人数</h3>
                        <p>{stats.maleCount}</p>
                    </div>
                </div>
                <div className="stat-card">
                    <FaFemale className="stat-icon" />
                    <div className="stat-info">
                        <h3>女生人数</h3>
                        <p>{stats.femaleCount}</p>
                    </div>
                </div>
                <div className="stat-card">
                    <FaChartBar className="stat-icon" />
                    <div className="stat-info">
                        <h3>平均年龄</h3>
                        <p>{stats.averageAge}</p>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default SystemOverview;
