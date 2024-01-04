import { useState, useEffect } from "react";
import { useRouter } from "next/router";
import axios from "axios";
import "@uiw/react-md-editor/markdown-editor.css";
import "@uiw/react-markdown-preview/markdown.css";
import dynamic from "next/dynamic";
import ShareDialog from "@/components/ShareDialog";

const MDEditor = dynamic(
  () => import("@uiw/react-md-editor").then((mod) => mod.default),
  { ssr: false }
);
// const EditorMarkdown = dynamic(
//   () =>
//     import("@uiw/react-md-editor").then((mod) => {
//       return mod.default.Markdown;
//     }),
//   { ssr: false }
// );

function DocumentEditor() {
  const router = useRouter();

  const [documentData, setDocumentData] = useState({
    data: null,
    loading: true,
    error: null,
  });

  const [text, setText] = useState("");

  console.log(router.query)

  const getDocument = async () => {
    try {
      const { data } = await axios.get(`/document/${router.query.id}`);
      setDocumentData({
        data,
        loading: false,
        error: null,
      });
      setText(data.body)
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
    <div className="flex flex-col w-full">
      <div className="max-w-7xl w-full mx-auto px-8 py-2 mb-6 flex justify-between align-center">
        <h1 className="font-medium uppercase text-3xl">{documentData.data?.title}</h1>
        <ShareDialog />
      </div>
      <div className="container">
        <MDEditor height={650} value={text} onChange={setText} />
      </div>
    </div>
  );
}

export default DocumentEditor;
