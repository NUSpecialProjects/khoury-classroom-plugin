import React from "react";
import "./styles.css";

interface IPanel {
  title: string;
  children: React.ReactNode;
  logo?: boolean;
}

const Panel: React.FC<IPanel> = ({ title, children, logo }) => {
  return (
    <div className="Panel">
      <div className="Panel__wrapper">
        {logo && (
          <div className="Panel__logo">
            <svg
              width="150"
              viewBox="0 0 525 102"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M77.7551 3.28228C77.0143 1.96241 75.7797 0.990755 74.3227 0.580951C73.6005 0.377591 72.8452 0.318625 72.1001 0.407424C71.355 0.496224 70.6348 0.731046 69.9805 1.09846C69.3263 1.46588 68.7509 1.95867 68.2873 2.54866C67.8237 3.13865 67.4809 3.81425 67.2787 4.53683C67.2787 4.53683 66.2268 7.68696 65.6642 9.90571L76.6642 13.7182L78.2729 7.61933C78.6821 6.16216 78.4958 4.60214 77.7551 3.28228Z"
                fill="#405DC5"
              />
              <path
                d="M75.1875 17.4995C75.1875 17.4995 55.125 92.484 54.5889 93.8429C54.0529 95.2019 48.0938 100.375 48.0938 100.375C46.375 101.843 45.7521 101.237 45.2197 99.2809C45.2197 99.2809 43.5956 93.4275 43.5947 90.7604C43.5938 88.0934 64.2812 13.7183 64.2812 13.7183L75.1875 17.4995Z"
                fill="#405DC5"
              />
              <path
                d="M34.3757 18.9838C35.4459 20.0543 36.047 21.506 36.047 23.0196C36.047 24.5333 35.4459 25.9849 34.3757 27.0554L14.1968 47.2401L34.3757 67.419C34.9061 67.949 35.3269 68.5783 35.614 69.2709C35.9012 69.9635 36.0492 70.7059 36.0494 71.4557C36.0497 72.2055 35.9023 72.948 35.6156 73.6408C35.3289 74.3336 34.9086 74.9631 34.3786 75.4935C33.8486 76.0238 33.2193 76.4446 32.5267 76.7318C31.8341 77.019 31.0917 77.1669 30.3419 77.1672C29.5922 77.1675 28.8497 77.02 28.1569 76.7334C27.4641 76.4467 26.8345 76.0263 26.3041 75.4964L2.08368 51.2759C1.01353 50.2054 0.412354 48.7537 0.412354 47.2401C0.412354 45.7264 1.01353 44.2748 2.08368 43.2043L26.3041 18.9838C27.3746 17.9137 28.8263 17.3125 30.3399 17.3125C31.8536 17.3125 33.3052 17.9137 34.3757 18.9838Z"
                fill="#405DC5"
              />
              <path
                d="M89.0958 27.0554C88.5506 26.5288 88.1157 25.899 87.8166 25.2025C87.5174 24.5061 87.3599 23.757 87.3533 22.9991C87.3467 22.2411 87.4912 21.4895 87.7782 20.7879C88.0652 20.0864 88.4891 19.4491 89.025 18.9131C89.561 18.3771 90.1984 17.9532 90.8999 17.6662C91.6014 17.3792 92.3531 17.2348 93.1111 17.2414C93.869 17.2479 94.618 17.4054 95.3145 17.7046C96.0109 18.0038 96.6408 18.4386 97.1674 18.9838L121.388 43.2043C122.458 44.2748 123.059 45.7264 123.059 47.2401C123.059 48.7537 122.458 50.2054 121.388 51.2759L97.1674 75.4964C96.0963 76.5667 94.6438 77.1677 93.1296 77.1672C91.6153 77.1667 90.1633 76.5646 89.0929 75.4935C88.0226 74.4224 87.4216 72.9699 87.4221 71.4557C87.4226 69.9414 88.0247 68.4894 89.0958 67.419L109.275 47.2401L89.0958 27.0554Z"
                fill="#405DC5"
              />
              <path
                d="M401.136 72.417L418.926 98.252H439.582L416.786 64.4182L437.498 33.1287H417.259L400.99 57.2374H393.448V2.0553H374.518V98.252H393.448V72.417H401.136Z"
                fill="#405DC5"
              />
              <path
                d="M463.153 99.5019C459.423 99.5019 455.791 99.2638 452.259 98.7875C448.767 98.3113 445.473 97.597 442.378 96.6446V82.7151C445.513 83.6675 448.827 84.4017 452.319 84.9176C455.811 85.3938 459.343 85.6319 462.915 85.6319C467.201 85.6319 470.217 85.0763 471.963 83.9652C473.749 82.8143 474.642 81.2269 474.642 79.2029C474.642 77.5759 474.205 76.3456 473.332 75.5122C472.459 74.6788 470.931 74.0439 468.749 73.6073L458.153 71.7025C452.041 70.5516 447.596 68.4284 444.819 65.333C442.041 62.2376 440.652 58.0905 440.652 52.8917C440.652 46.6612 443.152 41.641 448.152 37.8312C453.192 34.0215 460.911 32.1166 471.308 32.1166C474.483 32.1166 477.559 32.2952 480.535 32.6523C483.551 33.0095 486.27 33.5056 488.69 34.1405V48.07C486.151 47.3557 483.452 46.8199 480.595 46.4627C477.737 46.1056 474.82 45.927 471.844 45.927C468.233 45.927 465.455 46.2445 463.51 46.8794C461.566 47.4747 460.197 48.2684 459.403 49.2605C458.649 50.2527 458.272 51.3242 458.272 52.475C458.272 53.864 458.708 54.995 459.581 55.8681C460.454 56.7412 461.962 57.4158 464.105 57.8921L474.701 59.7969C480.495 60.8684 484.861 62.8527 487.797 65.7497C490.774 68.607 492.262 72.8732 492.262 78.5481C492.262 85.0168 489.722 90.1362 484.643 93.9063C479.603 97.6367 472.439 99.5019 463.153 99.5019Z"
                fill="#405DC5"
              />
              <path
                d="M512.107 99.2636C508.377 99.2636 505.38 98.2318 503.118 96.1682C500.896 94.1046 499.785 91.3068 499.785 87.7748C499.785 84.2031 500.876 81.3855 503.059 79.3219C505.281 77.2582 508.297 76.2264 512.107 76.2264C515.917 76.2264 518.913 77.2781 521.096 79.3814C523.318 81.4847 524.429 84.2825 524.429 87.7748C524.429 91.2274 523.298 94.0054 521.036 96.1087C518.814 98.212 515.837 99.2636 512.107 99.2636Z"
                fill="#405DC5"
              />
              <path
                d="M325.222 98.2518V33.1285H340.759L342.187 41.5219H343.14C344.965 38.3074 347.386 35.966 350.402 34.4976C353.418 32.9896 356.732 32.2356 360.343 32.2356C361.335 32.2356 362.288 32.2753 363.201 32.3547C364.113 32.434 364.927 32.5531 365.641 32.7118V49.8558C364.649 49.6574 363.578 49.5383 362.427 49.4986C361.315 49.4193 360.244 49.3796 359.212 49.3796C357.347 49.3796 355.462 49.6177 353.557 50.0939C351.652 50.5304 349.886 51.2249 348.259 52.1774C346.632 53.0901 345.263 54.2211 344.152 55.5704V98.2518H325.222Z"
                fill="#405DC5"
              />
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M260.921 94.2635C264.651 97.7161 269.85 99.4424 276.517 99.4424C280.049 99.4424 283.323 98.8272 286.339 97.597C289.355 96.3668 291.836 94.4222 293.78 91.7633H294.673L295.983 98.2518H311.579V59.4993C311.579 52.6338 310.527 47.2168 308.424 43.2482C306.321 39.2401 303.047 36.3827 298.602 34.6763C294.197 32.9698 288.542 32.1166 281.636 32.1166C278.382 32.1166 274.85 32.3547 271.041 32.8309C267.27 33.2674 263.798 33.9421 260.623 34.8549V49.1415C263.52 48.2287 266.556 47.5739 269.731 47.1771C272.906 46.7405 275.743 46.5223 278.243 46.5223C281.855 46.5223 284.732 46.8993 286.875 47.6533C289.058 48.3676 290.625 49.578 291.578 51.2845C292.57 52.9909 293.066 55.3324 293.066 58.3087V59.2193L282.47 60.0946C272.945 60.7295 266.04 62.7733 261.754 66.2259C257.508 69.6785 255.385 74.4804 255.385 80.6316C255.385 86.2272 257.23 90.7712 260.921 94.2635ZM293.066 69.8597V82.1198C291.558 83.6278 289.891 84.6993 288.065 85.3343C286.24 85.9296 284.454 86.2272 282.708 86.2272C279.89 86.2272 277.747 85.5923 276.279 84.3223C274.811 83.0127 274.076 81.1674 274.076 78.7862C274.076 76.3655 274.87 74.4606 276.458 73.0716C278.085 71.6826 280.724 70.8492 284.375 70.5714L293.066 69.8597Z"
                fill="#405DC5"
              />
              <path
                d="M148.059 98.2517V10.1508H171.573L196.111 65.7039L219.969 10.1508H243.482V98.2517H225.445V35.7774L203.718 85.5128H187.824L166.096 37.3491V98.2517H148.059Z"
                fill="#405DC5"
              />
            </svg>
          </div>
        )}
        <h1 className="Panel__title">{title}</h1>
        <div className="Panel__content">{children}</div>
      </div>
    </div>
  );
};

export default Panel;
