import React from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { buttonVariants } from "./ui/button";
import Listusers from "../components/ui/Listusers";

function ShareDialog() {
  return (
    <Dialog>
      <DialogTrigger className={buttonVariants()}>Share</DialogTrigger>

      <DialogContent>
        <DialogHeader>
          <DialogTitle>Share</DialogTitle>
        </DialogHeader>
        <DialogDescription>
          <Listusers />
        </DialogDescription>
      </DialogContent>
    </Dialog>
  );
}

export default ShareDialog;
