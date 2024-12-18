import React, { useState, useEffect, useCallback } from 'react';
import { studentApi } from '../api/studentApi';
import './EditStudent.css'; // 引入 CSS 文件

function EditStudent({ studentId, onStudentUpdated, onCancel }) {
    const [student, setStudent] = useState({
        id: '',
        name: '',
        age: '',
        gender: '',
        email: '',
        phone_number: ''
    });
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    const fetchStudent = useCallback(async () => {
        try {
            const response = await studentApi.getStudent(studentId);
            setStudent(response.data);
        } catch (err) {
            setError('无法获取学生数据');
            console.error(err);
        }
    }, [studentId]);

    useEffect(() => {
        fetchStudent();
    }, [fetchStudent]);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setStudent({
            ...student,
            [name]: name === 'age' ? parseInt(value, 10) : value
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setIsLoading(true);
        try {
            // 构建更新的学生对象，不包含不必要的字段
            const updatedStudent = {
                name: student.name,
                age: student.age,
                gender: student.gender,
                email: student.email,
                phone_number: student.phone_number
            };
            // 打印发送的数据用于调试
            console.log('更新学生数据：', updatedStudent);

            await studentApi.updateStudent(studentId, updatedStudent);
            setSuccess('学生信息更新成功');
            setError('');
            onStudentUpdated();
        } catch (err) {
            setError('更新学生信息失败');
            setSuccess('');
            console.error('更新学生时出错：', err.response ? err.response.data : err);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="edit-student-container">
            <h2>编辑学生信息</h2>
            {error && <p className="error-message">{error}</p>}
            {success && <p className="success-message">{success}</p>}
            <form onSubmit={handleSubmit} className="edit-student-form">
                <div className="form-group">
                    <label>ID:</label>
                    <input type="number" name="id" value={student.id} readOnly />
                </div>
                <div className="form-group">
                    <label>姓名:</label>
                    <input type="text" name="name" value={student.name} onChange={handleChange} required />
                </div>
                <div className="form-group">
                    <label>年龄:</label>
                    <input type="number" name="age" value={student.age} onChange={handleChange} required />
                </div>
                <div className="form-group">
                    <label>性别:</label>
                    <select name="gender" value={student.gender} onChange={handleChange} required>
                        <option value="">选择性别</option>
                        <option value="男">男</option>
                        <option value="女">女</option>
                    </select>
                </div>
                <div className="form-group">
                    <label>邮箱:</label>
                    <input type="email" name="email" value={student.email} onChange={handleChange} required />
                </div>
                <div className="form-group">
                    <label>电话:</label>
                    <input type="text" name="phone_number" value={student.phone_number} onChange={handleChange} required />
                </div>
                <div className="button-group">
                    <button type="submit" className="submit-button" disabled={isLoading}>
                        {isLoading ? '更新中...' : '更新学生'}
                    </button>
                    <button type="button" className="cancel-button" onClick={onCancel}>
                        取消
                    </button>
                </div>
            </form>
        </div>
    );
}

export default EditStudent;
