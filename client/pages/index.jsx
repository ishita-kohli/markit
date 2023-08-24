import { useState, useEffect, useContext } from "react";
import { API_URL } from "../constants";
import { AuthContext } from "../modules/auth_provider";
import { WebsocketContext } from "../modules/websocket_provider";
import { useRouter } from "next/router";
import axios from "axios";
import { Button } from "@/components/ui/button"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import DocumentList from "@/components/documentlist";




const Index = () => {
  const [documents, setDocuments] = useState([]);
  const [documentName, setDocumentName] = useState("");
  const { user } = useContext(AuthContext);
  const { setConn } = useContext(WebsocketContext);

  const router = useRouter();

  const getDocuments = async () => {
    try {
      throw new Error("not implemented");
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    getDocuments();
  }, []);

  const submitHandler = async (e) => {
    e.preventDefault();

    try {
      setDocumentName("");
      const { data } = await axios.post("/document", {
        title: documentName,
      });

      if (data) {
        console.log("works!");
      }
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <>
      <div className="my-8 px-4 md:mx-32 w-full h-full">
        <div className="flex justify-center mt-3 p-5">
          <input
            type="text"
            className="border border-gray-200 p-2 rounded-md focus:outline-none focus:border-blue-500 w-96"
            placeholder="document name"
            value={documentName}
            onChange={(e) => setDocumentName(e.target.value)}
          />
          <button
           className="bg-blue-500 border text-white rounded-md px-6 py-2 md:ml-4"
            onClick={submitHandler}
          >
            Create
          </button>
        </div>
        <hr class="h-px my-8 bg-gray-200 border-0 dark:bg-gray-700"></hr>
        <Tabs defaultValue="account" className="w-[400px]">
  <TabsList>
    <TabsTrigger value="owner">Owner</TabsTrigger>
    <TabsTrigger value="editor">Editor</TabsTrigger>
    <TabsTrigger value="viewer">Viewer</TabsTrigger>
  </TabsList>
  <TabsContent value="owner"><DocumentList role="owner" /></TabsContent>
  <TabsContent value="editor"><DocumentList role="editor" /></TabsContent>
  <TabsContent value="viewer"><DocumentList role="viewer" /></TabsContent>
</Tabs>
    
      </div>
    </>
  );
};

export default Index;
