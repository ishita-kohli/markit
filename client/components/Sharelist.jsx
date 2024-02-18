import React from "react";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import axios from "axios";
import { Avatar, AvatarFallback } from "./ui/avatar";
import { useState } from "react";
import { useEffect } from "react";

function Sharelist({ id, sharedWithRoles }) {
  const [userList, setUserList] = useState([]);

  const getUsers = async () => {
    try {
      const { data } = await axios.get("/users");
      console.log(data);
      setUserList(data);
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    getUsers();
  }, []);

  const users = userList.filter((user) =>
    sharedWithRoles.map((r) => r.user_id).includes(user.id)
  );
  return (
    <div className="flex -space-x-1">
      {users.map((user) => (
        <TooltipProvider key={user.id}>
          <Tooltip>
            <TooltipTrigger asChild>
              <Avatar className="outline-black outline">
                <AvatarFallback>
                  {user.username
                    .split(" ")
                    .map((s) => s[0].toUpperCase())
                    .join("")}
                </AvatarFallback>
              </Avatar>
            </TooltipTrigger>
            <TooltipContent>
              <p>{user.username} - ({sharedWithRoles.find(r => r.user_id === user.id).role})</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
      ))}
    </div>
  );
}

export default Sharelist;
