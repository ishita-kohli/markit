import { useState, useEffect, useContext } from "react";
import { API_URL } from "../constants";
import { AuthContext } from "../modules/auth_provider";
import { WebsocketContext } from "../modules/websocket_provider";
import { useRouter } from "next/router";
import axios from "axios";

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
            className="border border-grey p-2 rounded-md focus:outline-none focus:border-blue"
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
        <div className="mt-6">
          <div className="font-bold">Available Documents</div>
          <div className="grid grid-cols-1 md:grid-cols-5 gap-4 mt-6">
            {documents.map((document, index) => (
              <div
                key={index}
                className="border border-blue p-4 flex items-center rounded-md w-full"
              >
                <div className="w-full">
                  <div className="text-sm">document</div>
                  <div className="text-blue font-bold text-lg">
                    {document.name}
                  </div>
                </div>
                <div className="">
                  <button
                    className="px-4 text-white bg-blue rounded-md"
                    onClick={() => joinDocument(document.id)}
                  >
                    join
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
};

export default Index;
