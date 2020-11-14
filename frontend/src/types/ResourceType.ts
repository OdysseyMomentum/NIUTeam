
export interface IResourceType {
  uoi: string,
  resourceId: string,
  displayName: string,
  createdAt: number,
  resourceType: string,
  meta?: any
}

export interface IDocumentMeta {
  access?: string,
  filename: string,
  filesize: number,
  filetype: string
}
