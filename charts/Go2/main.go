package main

import "fmt"

// 1. Định nghĩa Struct (Giống như Controller của bạn)
type NhaHang struct {
	Ten     string
	SoKhach int
}

// -------------------------------------------------------------
// 2. HÀM KHỞI TẠO (Constructor)
// Dấu * chỉ định: Kết quả trả về là một ĐỊA CHỈ, không phải dữ liệu
// -------------------------------------------------------------
func NewNhaHang(ten string) *NhaHang {

	// Dấu & nghĩa là: Tạo một nhà hàng thật trong RAM,
	// sau đó lấy địa chỉ (chìa khóa) của nó để trả về.
	return &NhaHang{
		Ten:     ten,
		SoKhach: 0,
	}
}

// -------------------------------------------------------------
// 3. METHOD (Phương thức - Đã học ở bài trước)
// Dùng *NhaHang (Pointer Receiver) để thay đổi số khách thật
// -------------------------------------------------------------
func (n *NhaHang) DonKhachMoi() {
	n.SoKhach++
	fmt.Printf("[%s] Vừa đón thêm 1 khách. Tổng số khách: %d\n", n.Ten, n.SoKhach)
}

func main() {
	// Bước 1: Gọi hàm khởi tạo.
	// Biến 'quanPho' lúc này không chứa nguyên cái nhà hàng, nó chỉ cầm CHÌA KHÓA.
	quanPho := NewNhaHang("Phở Bát Đàn")

	// Thử cho một nhân viên A cầm chìa khóa này
	nhanVienA := quanPho

	// Thử cho nhân viên B cầm chung chìa khóa (địa chỉ)
	nhanVienB := quanPho

	// Bước 2: Cả 2 nhân viên đều gọi hàm đón khách
	fmt.Println("--- Bắt đầu kinh doanh ---")
	nhanVienA.DonKhachMoi() // Khách thứ 1
	nhanVienB.DonKhachMoi() // Khách thứ 2
	quanPho.DonKhachMoi()   // Ông chủ (quanPho) tự đón khách thứ 3

	// Bước 3: Kiểm tra kết quả
	fmt.Println("--------------------------")
	fmt.Printf("Chốt sổ: %s hôm nay có %d khách.\n", quanPho.Ten, quanPho.SoKhach)
}
