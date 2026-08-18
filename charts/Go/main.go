package main

import "fmt"

// Khai báo struct Ví Điện Tử
type ViDienTu struct {
	ChuTaiKhoan string
	SoDu        int
}

// --------------------------------------------------------
// 1. VALUE RECEIVER (Không có dấu *) - Lỗi logic nạp tiền
// --------------------------------------------------------
func (v ViDienTu) NapTienLoi(soTien int) {
	v.SoDu = v.SoDu + soTien
	fmt.Printf("[Bên trong hàm NapTienLoi] Số dư tạm thời: %d VNĐ\n", v.SoDu)
}

// --------------------------------------------------------
// 2. POINTER RECEIVER (Có dấu *) - Chuẩn logic nạp tiền
// --------------------------------------------------------
func (v *ViDienTu) NapTienChuan(soTien int) {
	v.SoDu = v.SoDu + soTien
	fmt.Printf("[Bên trong hàm NapTienChuan] Số dư thực tế: %d VNĐ\n", v.SoDu)
}

func main() {
	// Tạo một ví điện tử ban đầu có 100k
	viCuaToi := ViDienTu{
		ChuTaiKhoan: "Nguyen Van A",
		SoDu:        100000,
	}
	fmt.Println("Số dư ban đầu:", viCuaToi.SoDu, "VNĐ")
	fmt.Println("--------------------------------------------------")

	// Thử dùng hàm Value Receiver (Photocopy)
	viCuaToi.NapTienLoi(50000)
	fmt.Println("Sau khi gọi NapTienLoi, số dư gốc là:", viCuaToi.SoDu, "VNĐ")
	// -> KẾT QUẢ: Vẫn 100k! Tiền đã bị cộng vào bản sao và biến mất khi hàm kết thúc.

	fmt.Println("--------------------------------------------------")

	// Thử dùng hàm Pointer Receiver (Ghi trực tiếp)
	viCuaToi.NapTienChuan(50000)
	fmt.Println("Sau khi gọi NapTienChuan, số dư gốc là:", viCuaToi.SoDu, "VNĐ")
	// -> KẾT QUẢ: 150k! Dữ liệu gốc đã được cập nhật thành công.
}
