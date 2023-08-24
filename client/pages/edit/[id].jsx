import React from 'react'
import { useRouter } from 'next/router'


function DocumentEditor() {
  const router = useRouter()
  return (
    <div>
      Document - {router.query.id}
    </div>
  )
}

export default DocumentEditor
