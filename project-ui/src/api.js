import axios from 'axios'

const url = 'http://localhost:8080/api/v1/'

const instance = axios.create({
    baseURL: url,
    headers: {
      'Accept' : 'application/json',
      'Content-Type': 'application/json',
      }
  });

function getUsers() {
    return instance
        .get(url+'users')
        .then(response => response.data)
        .catch((error) => {
            console.log("Error when getting users: " + error);
        })
}

function updateUser(id, user) {
    return instance.put(`${url}users/${id}`, {
        nickname: user.nickname,
        email: user.email,
        password: user.password,
        company_id: parseInt(user.company_id)
    }).catch((error) => {
        console.log("Error while updating user: " + error);
    })
}

function deleteUser(id) {
    return instance.delete(`${url}users/${id}`).catch((error) => {
        console.log("Error while deleting user: " + error);
    })
}

function createUser(user) {
    console.log(user);
    return instance.post('users', {
        nickname: user.nickname,
        email: user.email,
        password: user.password,
        company_id: parseInt(user.company_id)
    }).catch((error) => {
        console.log("Error while creating user: " + error);
    })
}

function getCompanies() {
    return instance
        .get(url+'companies')
        .then(response => response.data)
        .catch((error) => {
            console.log("Error when getting companies: " + error);
        })
}

function updateCompany(id, company) {
    return instance.put(`${url}companies/${id}`, {
        name: company.name
    }).catch((error) => {
        console.log("Error updating updating company: " + error);
    })
}

function deleteCompany(id) {
    return instance.delete(`${url}companies/${id}`).catch((error) => {
        console.log("Error while deleting company: " + error);
    })
}

function createCompany(company) {
    return instance.post('companies', {
        name: company.name
    }).catch((error) => {
        console.log("Error while creating company: " + error);
    })
}

export default {
    getUsers,
    createUser,
    updateUser,
    deleteUser,
    getCompanies,
    updateCompany,
    deleteCompany,
    createCompany,
}