import React, { useState } from "react";
import { useEffect } from "react";
import axios from "axios";
import { cn } from "@/lib/utils";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Button } from "./button";

const Listusers = ({id}) => {
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
  const submitHandler = async (e) => {
    e.preventDefault();

    try {
       
      const { data } = await axios.patch(`/document/${id}/share`, {
        share_with :selectedUser.id,
        role: selectedRole,
      });

      if (data) {
        console.log("works!");
      }
    } catch (err) {
      console.log(err);
    }
  };


  return (
    <div>
      <ul>
        <div className="font-bold text-gray-800 grid grid-cols-[1fr_2fr] py-2">
          <div>Username</div>
          <div>E-mail</div>
        </div>
        {users?.map((user) => (
          <li
            key={user.id}
            onClick={() => handleUserSelect(user)}
            className={cn(
              "grid grid-cols-[1fr_2fr_1fr] py-1 border-b border-stone-500",
              selectedUser === user ? "bg-stone-50" : ""
            )}
          >
            <div>{user.username}</div>
            <div>{user.email}</div>
            <div>
              {selectedUser === user ? (
                <Select value={selectedRole} onValueChange={handleRoleSelect}>
                  <SelectTrigger>
                    <SelectValue placeholder="Role" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="editor">Editor</SelectItem>
                    <SelectItem value="viewer">Viewer</SelectItem>
                  </SelectContent>
                </Select>
              ) : null}
            </div>
          </li>
        ))}
      </ul>

      <Button onClick={submitHandler}>Save</Button>
    </div>
  );
};

export default Listusers;
