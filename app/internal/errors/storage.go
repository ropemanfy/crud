package errors

var (
	ErrFilestorage          = new("filestorage error", "errorFS")
	ErrS3CreateBucket       = new("create s3 bucket error", "s3.client.createBucket")
	ErrS3Unexpected         = new("unexpected error occurred", "s3.client.unexpected")
	ErrS3IsBucketExists     = new("is bucket exists error", "s3.client.isBucketExists")
	ErrS3DeleteBucket       = new("delete bucket error", "s3.client.deleteBucket")
	ErrS3PutObject          = new("put object error", "s3.client.putObject")
	ErrS3PutObjectMultipart = new("multipart put object error", "s3.client.putObjectMultipart")
	ErrS3GetObject          = new("get object error", "s3.client.getObject")
	ErrS3ListBucketObject   = new("list bucket object error", "s3.client.listBucketObject")
	ErrS3DeleteObject       = new("delete object error", "s3.client.deleteObject")
	ErrS3IsObjectExists     = new("is object exists error", "s3.client.isObjectExists")
)
