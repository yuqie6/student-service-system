import axios from 'axios';
import { apiConfig } from './config';

const api = axios.create(apiConfig);

export const studentApi = {
    getAllStudents: () => api.get('/students'),
    getStudent: (id) => api.get(`/students/${id}`),
    createStudent: (student) => api.post('/students', student),
    updateStudent: (id, student) => api.put(`/students/${id}`, student),
    deleteStudent: (id) => api.delete(`/students/${id}`)
};
