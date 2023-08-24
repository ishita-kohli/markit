import { useState, useEffect } from "react";
import { useRouter } from "next/router";
import axios from "axios"
import MarkdownProvider from "@/providers/markdown";

function DocumentEditor() {
  const router = useRouter();

  const [documentData, setDocumentData] = useState({
    data: null,
    loading: true,
    error: null,
  });

  const getDocument = async () => {
    try {
      const { data } = await axios.get(`/document/${router.query.id}`);
      setDocumentData({
        data,
        loading: false,
        error: null,
      });
    } catch (err) {
      setDocumentData({
        error: err,
        loading: false,
        data: null,
      });
    }
  };

  useEffect(() => {
    getDocument();
  }, []);

  return (
    <MarkdownProvider>
      {documentData.data && (
        <>
          <div className="flex">
            <h1>{documentData.data.title}</h1>
          </div>
          <div className="grid grid-cols-2 w-full h-screen">
            <div className="bg-red-300"></div>
            <div className="bg-green-300"></div>
          </div>
        </>
      )}
      {documentData.loading && <p>Loading...</p>}
      {documentData.error && <p>Failed to load data: {JSON.stringify(documentData)}</p>}
    </MarkdownProvider>
  );
}

export default DocumentEditor;
