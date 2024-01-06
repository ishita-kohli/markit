import React, { useState } from "react";
import {useEffect} from "react";
import axios from "axios";

const Listusers = () => {
    const [users, setUsers] = useState(null);
  const [selectedUser, setSelectedUser] = useState(null);
  const [selectedRole, setSelectedRole] = useState(null);
  const getUsers = async () => {
    try {
      const { data } = await axios.get("/users");
      console.log(data);
      setUsers(data);
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    getUsers();
  }, []);
  const handleUserSelect = (user) => {
    setSelectedUser(user);
    setSelectedRole(null); 
  };

  const handleRoleSelect = (role) => {
    setSelectedRole(role);
  };

  return (
    <div>
      
      <ul>
      <div className= "bg-purple-200 text-gray-800 font-bold columns-2 ">
        <div className>Username</div>
        <div className>E-mail</div>
        
    </div>
        {users?.map(user => (
          <li key={user.id} onClick={() => handleUserSelect(user)}>
           
           <div className=" columns-2 ">
        <div className>{JSON.stringify(user.username)}</div>
        <div className>{JSON.stringify(user.email)}</div>
        
    </div>
         <hr/>   

          </li>
        ))}
      </ul>

      {/*selectedUser && (
        <div>
          <label >Select Role:</label>
          <select
            id="roleSelect"
            onChange={(e) => handleRoleSelect(e.target.value)}
            value={selectedRole || ''}
          >
            <option value="" disabled>Select Role</option>
            <option value="editor">Editor</option>
            <option value="viewer">Viewer</option>
          </select>
          {selectedRole && (
            <p>{`User ${selectedUser} is set as ${selectedRole}`}</p>
          )}
        </div>
          )*/}
    </div>
  );
};

export default Listusers;
