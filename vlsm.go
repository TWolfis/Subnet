/*
use this program to find out information about network addressing
input an IP address with its network bits in a addr/bits format 
find out the network address, broadcast address leading bits, mask length and maximum number of hosts
that is supported by the network provided
provide as args as many valid networks as you like

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

Author: Thomas Wolfis
*/
package main

import (
	"fmt"
	"net"
	"os"
	"encoding/binary"
	"math"
	"time"
)

func main(){
	if len(os.Args) <= 1{
		fmt.Fprintf(os.Stderr, "Usage: %s  address/network-bits \nFor example 192.168.1.0/24\n",os.Args[0])
		os.Exit(1)
	}
	start := time.Now()
	failed := 0

	//print the column names
	fmt.Printf("Addr\t\tNetAddr\t\tMask(hex)\tBrdAddr\t\tLeading ones\tMask Length\tMax Hosts\n")
	
	//itterate over all addresses given as input
	for i :=1; i < len(os.Args);i++{

		//parse the given address
		ip,network,err := net.ParseCIDR(os.Args[i])
		//check if the address contains any errors

		if err == nil{
			//continue with getting information gathering on valid networks
			ones,bits := network.Mask.Size()
			max_hosts := math.Pow(2,32-float64(ones))-2
			
			//calculate the broadcast address by converting the net and mask to binary and performing bitwise operations on them
			bin_net := binary.BigEndian.Uint32(network.IP)
			bin_mask := binary.BigEndian.Uint32(network.Mask)
			brd := (bin_net & bin_mask)| (bin_mask ^ 0xffffffff)

			//fill in an ip address struct with the bytes stored in brd
			brdAddr := make(net.IP,4)
			binary.BigEndian.PutUint32(brdAddr,brd)
			
			//print out the values
			fmt.Printf("%v\t%v\t%v\t%v\t%v\t\t%v\t\t%v\n",ip,network.IP.String(),network.Mask.String(),brdAddr.String(),ones,bits,max_hosts)
		}else{
			//error handling
			fmt.Println(err)
			failed++
		}

	}
	fmt.Printf("Processed %d networks in %.3fs of which %d failed\n",len(os.Args)-1,time.Since(start).Seconds(),failed)

}
