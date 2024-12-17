import React, { useState } from 'react';
import { studentApi } from '../api/studentApi';
import { IoIosRocket } from 'react-icons/io'; // 引入图标库
import './AddStudent.css'; // 引入CSS文件

function AddStudent({ onStudentAdded }) {
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
    const [errors, setErrors] = useState({});

    const handleChange = (e) => {
        const { name, value } = e.target;
        setStudent({
            ...student,
            [name]: name === 'id' || name === 'age' ? parseInt(value, 10) : value
        });
    };

    const validate = () => {
        const newErrors = {};
        if (!student.name) newErrors.name = '姓名不能为空';
        if (!student.age) newErrors.age = '年龄不能为空';
        if (!student.gender) newErrors.gender = '请选择性别';
        if (!student.email) newErrors.email = '邮箱不能为空';
        if (!student.phone_number) newErrors.phone_number = '电话不能为空';
        setErrors(newErrors);
        return Object.keys(newErrors).length === 0;
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!validate()) return;
        setIsLoading(true);
        try {
            // 打印数据类型以进行调试
            console.log('Submitting student:', student);
            console.log('Data types:', {
                id: typeof student.id,
                name: typeof student.name,
                age: typeof student.age,
                gender: typeof student.gender,
                email: typeof student.email,
                phone_number: typeof student.phone_number
            });

            await studentApi.createStudent(student);
            setSuccess('学生添加成功');
            setError('');
            setStudent({
                id: '',
                name: '',
                age: '',
                gender: '',
                email: '',
                phone_number: ''
            });
            onStudentAdded();
        } catch (err) {
            setError('添加学生失败');
            setSuccess('');
            console.error('Error adding student:', err.response ? err.response.data : err); // 添加调试信息
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="add-student-container fade-in-up">
            <h2>学生信息录入</h2>
            {error && <p className="error-message">{error}</p>}
            {success && <p className="success-message">{success}</p>}
            <form onSubmit={handleSubmit} className="add-student-form">
                <div className="form-group">
                    <label>ID:</label>
                    <input type="number" name="id" value={student.id} onChange={handleChange} required />
                    {errors.id && <p className="input-error">{errors.id}</p>}
                </div>
                <div className="form-group">
                    <label>姓名:</label>
                    <input type="text" name="name" value={student.name} onChange={handleChange} required />
                    {errors.name && <p className="input-error">{errors.name}</p>}
                </div>
                <div className="form-group">
                    <label>年龄:</label>
                    <input type="number" name="age" value={student.age} onChange={handleChange} required />
                    {errors.age && <p className="input-error">{errors.age}</p>}
                </div>
                <div className="form-group">
                    <label>性别:</label>
                    <select name="gender" value={student.gender} onChange={handleChange} required>
                        <option value="">选择性别</option>
                        <option value="男">男</option>
                        <option value="女">女</option>
                    </select>
                    {errors.gender && <p className="input-error">{errors.gender}</p>}
                </div>
                <div className="form-group">
                    <label>邮箱:</label>
                    <input type="email" name="email" value={student.email} onChange={handleChange} required />
                    {errors.email && <p className="input-error">{errors.email}</p>}
                </div>
                <div className="form-group">
                    <label>电话:</label>
                    <input type="text" name="phone_number" value={student.phone_number} onChange={handleChange} required />
                    {errors.phone_number && <p className="input-error">{errors.phone_number}</p>}
                </div>
                <button type="submit" className="submit-button" disabled={isLoading}>
                    {isLoading ? '提交中...' : <>
                        <IoIosRocket className="button-icon" />
                        添加学生
                    </>}
                </button>
            </form>
        </div>
    );
}

export default AddStudent;
