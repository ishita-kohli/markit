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
  const id = router.query.id;

  const [documentData, setDocumentData] = useState({
    data: null,
    loading: true,
    error: null,
  });

  const [text, setText] = useState("");

  const getDocument = async () => {
    if (!id) return;
    try {
      const { data } = await axios.get(`/document/${id}`);
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
  }, [id]);

  return (
    <div className="flex flex-col w-full">
      <div className="flex justify-between w-full px-8 py-2 mx-auto mb-6 max-w-7xl align-center">
        <h1 className="text-3xl font-medium uppercase">{documentData.data?.title}</h1>
        <ShareDialog />
      </div>
      <div className="container">
        <MDEditor height={650} value={text} onChange={setText} />
      </div>
    </div>
  );
}

export default DocumentEditor;
