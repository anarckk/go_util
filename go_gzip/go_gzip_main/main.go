package main

import "gitea.bee.anarckk.me/anarckk/go_util/go_gzip"

//lint:ignore U1000 Ignore unused function temporarily for debugging
func testCompress() {
	err := go_gzip.CompressFolder("/pet-workdir/data2/pet-dev-var/var/pet-workspace/io-test-space/minio", "/pet-workdir/data2/pet-dev-var/var/pet-workspace/io-test-space/go_gzip/minio.tar.gz")
	if err != nil {
		panic(err)
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func testDepress() {
	go_gzip.DeCompress("/pet-workdir/data2/pet-dev-var/var/pet-workspace/io-test-space/go_gzip/minio.tar.gz", "/pet-workdir/data2/pet-dev-var/var/pet-workspace/io-test-space/go_gzip")
}
func main() {
	testDepress()
}
