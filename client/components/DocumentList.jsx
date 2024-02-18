import { useState, useEffect } from "react";
import Link from "next/link";
import axios from "axios";
import { Card, CardDescription, CardFooter, CardHeader, CardTitle } from "./ui/card";
import { buttonVariants } from "./ui/button";

function DocumentList({ role }) {
  const [documents, setDocuments] = useState([]);

  const getDocuments = async () => {
    try {
      const { data } = await axios.get("/document", {
        params: {
          role,
        },
      });
      setDocuments(data);
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    getDocuments();
  }, []);

  return (
    <div className="grid gap-4 grid-cols-auto">
      {documents.map((document) => (
          <Card key={document.id}>
            <CardHeader>
              <CardTitle>{document.title}</CardTitle>
              <CardDescription>{role}</CardDescription>
            </CardHeader>
            <CardDescription>
              {document.body}
            </CardDescription>
            <CardFooter>
              <Link href={`edit/${document.id}`} className={buttonVariants()}>Open &rarr;</Link>
            </CardFooter>
          </Card>
))}
    </div>
  );
}

export default DocumentList;
