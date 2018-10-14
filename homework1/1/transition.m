function transition()

img1 = imread('Åµ±´¶û.jpg');
size1 = size(img1);
len = size1(1,1);
hei = size1(1,2);
img2 = imread('lena.jpg');
radius = 0;
point = zeros(1,2);
point(1,1) = len / 2;
point(1,2) = hei / 2;
max = sqrt((len / 2)^2 + (hei/2)^2);
disp(max);
while(radius <= max)
    for i = 1 : len
        for j = 1 : hei
            if((i-point(1,1))^2 + (j-point(1,2))^2 <= radius * radius)
                img1(i,j,:) = img2(i,j,:);
            end
        end
    end
    imshow(img1);
    pause(0.05);
    radius = radius + 1;
end