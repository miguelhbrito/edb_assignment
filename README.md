# edb_assignment

## Problem statement

Grandmother has an old cell phone that can only take pictures in PBM format. She is wanting you to write an application in GO for her that will rotate a given arbitrary PBM image 90 degrees clockwise and output the result. Grandmother also prefers that we work from first principles and do not use any third party libraries to help. If possible, Grandmother has a feature request, and would like to be able to rotate her PBM images clockwise or counterclockwise by an arbitrary number of degrees.
PBM

PBM is an image format that can be represented in ascii and is easily manipulated.

an example pbm file could look like this

P1

## This is an example bitmap of the letter "J"

6 10

0 0 0 0 1 0

0 0 0 0 1 0

0 0 0 0 1 0

0 0 0 0 1 0

0 0 0 0 1 0

0 0 0 0 1 0

1 0 0 0 1 0

0 1 1 1 0 0

0 0 0 0 0 0

0 0 0 0 0 0

 

We can assume that Grandmother's phone uses pbm files with the magic number P1 only. Alas, grandmother's phone doesn't always generate square images, it can create rectangular images as well.

## 🧰 How to use
Generate the bin file:
```powershell
go build cmd/grandmotherapp.go
```
And then:
```powershell
./grandmotherapp testImage 90
```
Or just type:
```powershell
go run cmd/grandmotherapp.go testImage 90
```
The image can be rotated for with this degres:
```powershell
90, 180, 270, -90, -180, -270 and reverse
```
